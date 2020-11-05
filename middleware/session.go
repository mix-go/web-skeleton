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

        _, ok := store.Get("user_id")
        if !ok {
            c.Status(401)
            c.Abort()
            return
        }

        c.Next()
    }
}
