package chatgpt

import (
	"context"
	"testing"
)

func TestClient_Chat(t *testing.T) {
	client := NewClient(WithApiKey("sk-xxxxxx"))
	chatRequest := ChatCompletionRequestBody{
		Model: "gpt-3.5-turbo",
		Messages: []ChatMessage{
			{Role: "user", Content: "who is golang author?"},
		},
	}
	resp, err := client.Chat(context.Background(), chatRequest)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
