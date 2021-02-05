package server

import (
	"github.com/THEToilet/events-server/pkg/server/handler"
	"github.com/THEToilet/events-server/pkg/usercase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func NewServer(userUseCase *usercase.UserUseCase, eventUseCase *usercase.EventUseCase, tagUseCase *usercase.TagUseCase, authUseCase *usercase.AuthUseCase) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userHandler := handler.NewUserHandler(userUseCase)
	tagHandler := handler.NewTagHandler(tagUseCase)
	eventHandler := handler.NewEventHandler(eventUseCase)
	authHandler := handler.NewAuthHandler(authUseCase)

	v1 := e.Group("/api/v1")
	auth := v1.Group("/auth")
	auth.GET("/callback", authHandler.CallBack)
	auth.GET("/login", authHandler.Login)

	users := v1.Group("/users")
	users.GET("", userHandler.GetUser)
	users.POST("/login", userHandler.GetUser)
	users.POST("/entry", userHandler.GetUser)
	users.GET("/logout", userHandler.GetUser)
	events := v1.Group("/events")
	events.GET("", eventHandler.GetEvent)
	events.POST("", eventHandler.PostEvent)
	events.PUT("/:id", eventHandler.PutEvent)
	events.DELETE("/:id", eventHandler.DeleteEvent)

	tags := events.Group("/tags")
	tags.GET("", tagHandler.GetTagList)
	tags.POST("", tagHandler.PostTagList)
	return e
}

/*
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
 */
}
