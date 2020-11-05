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
    db := globals.DB()
    if err := db.Create(&models.User{
        Name:     c.Request.PostFormValue("name"),
        CreateAt: time.Now(),
    }).Error; err != nil {
        c.JSON(http.StatusOK, gin.H{
            "status":  http.StatusInternalServerError,
            "message": err.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "status":  http.StatusOK,
        "message": "ok",
    })
}

func (t *UserController) Save(c *gin.Context) {
    db := globals.DB()
    if err := db.Create(&models.User{
        Name:     c.Request.PostFormValue("name"),
        CreateAt: time.Now(),
    }).Error; err != nil {
        c.JSON(http.StatusOK, gin.H{
            "status":  http.StatusInternalServerError,
            "message": err.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "status":  http.StatusOK,
        "message": "ok",
    })
}
