package handler

import (
	"github.com/THEToilet/events-server/pkg/http"
	"fmt"
	"github.com/THEToilet/events-server/pkg/http/response"
	"github.com/THEToilet/events-server/pkg/usercase"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler struct {
	authUseCase *usercase.AuthUseCase
}

func (h AuthHandler) CallBack(context echo.Context) error {
	
}

func (h AuthHandler) Login(context echo.Context) error {
	
}

func NewAuthHandler(authUseCase *usercase.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUseCase: authUseCase,
	}
}

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