package chatgpt

type Option interface {
	apply(*Client)
}

type apiKeyOption struct {
	apiKey string
}

func (a apiKeyOption) apply(client *Client) {
	client.APIKey = a.apiKey
}

// WithApiKey Set client api key
func WithApiKey(apikey string) Option {
	return apiKeyOption{apiKey: apikey}
}

type baseUrlOption struct {
	baseUrl string
}

func (b baseUrlOption) apply(client *Client) {
	client.BaseURL = b.baseUrl
}

// WithBaseUrl Set client base url
func WithBaseUrl(baseUrl string) Option {
	return baseUrlOption{baseUrl: baseUrl}
}
