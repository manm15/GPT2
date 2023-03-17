package chatgpt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client *Client) buildRequest(ctx context.Context, method string, body interface{}) (req *http.Request, err error) {
	r, contentType, err := client.bodyToReader(body)
	if err != nil {
		return nil, fmt.Errorf("failed to buildRequest request buf from given body: %v", err)
	}
	req, err = http.NewRequest(method, client.BaseURL, r)
	if err != nil {
		return nil, fmt.Errorf("failed to init request: %v", err)
	}
	req.Header.Add("Content-Type", contentType)
	if client.APIKey != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.APIKey))
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}
	return req, nil
}

func (client *Client) execute(req *http.Request, response interface{}) error {
	if client.HttpClient == nil {
		client.HttpClient = http.DefaultClient
	}
	httpResp, err := client.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode >= 400 {
		return client.apiError(httpResp)
	}
	if err := json.NewDecoder(httpResp.Body).Decode(response); err != nil {
		return fmt.Errorf("failed to decode response to %T: %v", response, err)
	}
	return nil
}

func (client *Client) bodyToReader(body interface{}) (io.Reader, string, error) {
	var r io.Reader
	switch v := body.(type) {
	// case io.Reader:
	// 	r = v
	case nil:
		r = nil
	case MultipartFormDataRequestBody: // TODO: Refactor
		buf, ct, err := v.ToMultipartFormData()
		if err != nil {
			return nil, "", err
		}
		return buf, ct, nil
	default:
		b, err := json.Marshal(body)
		if err != nil {
			return nil, "", err
		}
		r = bytes.NewBuffer(b)
	}
	return r, "application/json", nil
}

func call[T any](ctx context.Context, client *Client, method string, p string, body interface{}, resp T) (T, error) {
	req, err := client.buildRequest(ctx, method, body)
	if err != nil {
		return resp, err
	}
	err = client.execute(req, &resp)
	return resp, err
}
