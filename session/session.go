package session

import (
	resty "gopkg.in/resty.v1"
)

type SessionRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Response struct {
	ApiVersion string `json:"api-version"`
	Context    string `json:"context"`
}

type SessionResponseData struct {
	User         *SessionUser `json:"user,omitempty"`
	SessionToken string       `json:"session-token"`
}

type SessionUser struct {
	Email      string `json:"email"`
	Username   string `json:"username"`
	ExternalId string `json:"external-id"`
}

type SessionResponse struct {
	Data *SessionResponseData `json:"data,omitempty"`
}

func Login(login string, password string) (*resty.Response, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept-Version", "v1").
		SetBody(&SessionRequest{
			Login:    login,
			Password: password,
		}).
		SetResult(&SessionResponse{}).
		Post("https://api.tastyworks.com/sessions")

	return resp, err
}
