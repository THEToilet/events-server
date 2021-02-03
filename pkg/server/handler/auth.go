package handler

import (
	"../../http/response"
	"github.com/THEToilet/events-server/pkg/http"
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

func HandleUserLogin() http.HandlerFunc {
	err := request.ParseForm()
	user, err := data.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
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
