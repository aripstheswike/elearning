package controller

import (
	"elearning/lib"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Dashboard(c *gin.Context) {
	lib.Json(c, http.StatusOK, "Selamat datang di Pembelajaran Online", gin.H{})
	return
}
