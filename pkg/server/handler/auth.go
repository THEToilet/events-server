package handler

import (
	"../../http/response"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

func HandleAuthLogin() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		baseURL := "http://localhost:3000/oauth2/v1/authorize"
		redirectURL := "http://localhost:4000/callback"
		scope := "openid profile"
		nonce, _ := uuid.NewUUID()
		state, _ := uuid.NewUUID()
		fmt.Println(redirectURL + scope + nonce.String() + state.String() + baseURL)
		response.Success(writer, nil)
	}
}

type params struct {
	Code        string `json:"code"`
	ClientID    string `json:"client_id"`
	RedirectURL string `json:"redirect_url"`
	State       string `json:"state"`
	Scope       string `json:"scope"`
	Nonce       string `json:"nonce"`
}
