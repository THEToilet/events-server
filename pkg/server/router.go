package main

import (
	"github.com/THEToilet/events-server/pkg/server/handler"
	"github.com/THEToilet/events-server/pkg/usercase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func NewServer(userUseCase *usercase.UserUseCase) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	userHandler := handler.NewUserHandler(userUseCase)
	v1 := e.Group("/api/v1")
	v1.GET("/callback")

	users := v1.Group("/users")
	users.GET("", userHandler.GetUser)
	users.POST("/login", userHandler.GetUser)
	users.POST("/entry", userHandler.GetUser)
	users.GET("/logout", userHandler.GetUser)
	events := v1.Group("/events")
	events.GET("")
	events.POST("")
	events.PUT("/:id")
	events.DELETE("/:id")

	tags := events.Group("/tags")
	tags.GET("")
	tags.POST("")
	return e
}

// get GETリクエストを処理する
func get(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodGet)
}

// post POSTリクエストを処理する
func post(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodPost)
}

// post PUTリクエストを処理する
func put(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodPut)
}

// post DELETEリクエストを処理する
func delete(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodDelete)
}

// httpMethod 指定したHTTPメソッドでAPIの処理を実行する
func httpMethod(apiFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// CORS対応
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin,x-token")

		// プリフライトリクエストは処理を通さない
		if request.Method == http.MethodOptions {
			return
		}
		// 指定のHTTPメソッドでない場合はエラー
		if request.Method != method {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Method Not Allowed"))
			return
		}

		// 共通のレスポンスヘッダを設定
		writer.Header().Add("Content-Type", "application/json")
		apiFunc(writer, request)
	}
}
