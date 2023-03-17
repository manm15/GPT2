package chatgpt

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponseBody struct {
	Error ErrorEntry `json:"error"`
}

type ErrorEntry struct {
	Message string      `json:"message"`
	Type    string      `json:"type"`
	Param   interface{} `json:"param"` // TODO: typing
	Code    interface{} `json:"code"`  // TODO: typing
}

func (err ErrorEntry) Error() string {
	return fmt.Sprintf("%v: %v (param: %v, code: %v)", err.Type, err.Message, err.Param, err.Code)
}

func (client *Client) apiError(res *http.Response) error {
	errBody := ErrorResponseBody{}
	if err := json.NewDecoder(res.Body).Decode(&errBody); err != nil {
		return fmt.Errorf("failed to decode error body: %v", err)
	}
	return fmt.Errorf("openai api error: %v", errBody.Error)
}
