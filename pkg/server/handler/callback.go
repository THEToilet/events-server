package handler

import (
	"net/http"
)

func HandleCallback() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
	}
}

type data struct {
	GrantType string `json:"authorization_code"`
	Code string `json:"code"`
	RedirectURI string `json:redirect_uri`
	ClientID string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}