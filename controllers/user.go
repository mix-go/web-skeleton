package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/mix-go/web-skeleton/globals"
    "github.com/mix-go/web-skeleton/models"
    "net/http"
    "time"
)

type UserController struct {
}

func (t *UserController) Add(c *gin.Context) {
    // 网页
    if c.Request.Method == http.MethodGet {
        c.HTML(http.StatusOK, "user_add.tmpl", gin.H{
            "title": "User add",
        })
        c.Abort()
        return
    }

    db := globals.DB()
    if err := db.Create(&models.User{
        Name:     c.Request.PostFormValue("name"),
        CreateAt: time.Now(),
    }).Error; err != nil {
        panic(err)
    }
    c.String(http.StatusOK, "Add ok!")
}
