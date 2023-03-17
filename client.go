package chatgpt

import (
	"context"
	"net/http"
)

const (
	defaultApiBaseURL = "https://chatgpt-api.shn.hk/v1/"
)

type Client struct {
	// APIKey issued by OpenAI console.
	// See https://beta.openai.com/account/api-keys
	APIKey string

	// BaseURL of API including the version.
	// e.g., https://chatgpt-api.shn.hk/v1/
	BaseURL    string
	HttpClient *http.Client
}

func NewClient(options ...Option) *Client {
	client := &Client{BaseURL: defaultApiBaseURL}
	for _, opt := range options {
		opt.apply(client)
	}
	return client
}

func (client *Client) Chat(ctx context.Context, body ChatCompletionRequestBody) (resp ChatCompletionResponse, err error) {
	return call(ctx, client, http.MethodPost, client.BaseURL, body, resp)
}
