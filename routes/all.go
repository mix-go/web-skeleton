package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/mix-go/web-skeleton/controllers"
    "github.com/mix-go/web-skeleton/middleware"
)

var RouteDefinitionCallbacks []func(router *gin.Engine)

func init() {
    RouteDefinitionCallbacks = append(RouteDefinitionCallbacks, func(router *gin.Engine) {
        router.GET("hello",
            func(ctx *gin.Context) {
                hello := controllers.HelloController{}
                hello.Index(ctx)
            },
        )

        router.Any("users/add",
            middleware.SessionMiddleware(),
            func(ctx *gin.Context) {
                hello := controllers.UserController{}
                hello.Add(ctx)
            },
        )

        router.Any("login", func(ctx *gin.Context) {
            auth := controllers.LoginController{}
            auth.Index(ctx)
        })
    })
}
