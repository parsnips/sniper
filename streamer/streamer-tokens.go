package streamer

import (
	resty "gopkg.in/resty.v1"
)

type Response struct {
	Context string `json:"context"`
}

type StreamerTokenResponse struct {
	Data *StreamerTokenResponseData `json:"data,omitempty"`
}

type StreamerTokenResponseData struct {
	Token        string `json:"token"`
	StreamerUrl  string `json:"streamer-url"`
	WebsocketUrl string `json:"websocket-url"`
	Level        string `json:"level"`
}

func GetStreamerToken(authToken string) (*resty.Response, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept-Version", "v1").
		SetHeader("Authorization", authToken).
		Get("https://api.tastyworks.com/quote-streamer-tokens")

	return resp, err
}
