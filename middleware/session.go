package middleware

import (
    "context"
    "github.com/gin-gonic/gin"
    "github.com/go-session/session"
)

func SessionMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        store, err := session.Start(context.Background(), c.Writer, c.Request)
        if err != nil {
            panic(err)
        }
        if _, ok := store.Get("userinfo"); !ok {
            c.Status(401)
            c.Abort()
            return
        }

        c.Next()
    }
}
