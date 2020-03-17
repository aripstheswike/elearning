package lib

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var MapKosong map[string]interface{}
var SliceMapKosong map[string]interface{}

func Json(c *gin.Context, code, msg, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func Pages(c *gin.Context, code, msg interface{}, title, template string, data map[string]interface{}, js, css []string) {
	c.HTML(http.StatusOK, template, gin.H{
		"title": title,
		"msg":   msg,
		"css":   css,
		"js":    js,
		"data":  data,
	})
}

func JsonError(c *gin.Context, err error){
	Json(c, http.StatusInternalServerError, "Terjadi Kesalahan-"+err.Error(), gin.H{})
	return
}

func GetDataJwt(c *gin.Context) (map[string]interface{}, error) {
	data, exist := c.Get("jwt")
	if !exist{
	return nil, errors.New("Data Dari Token Tidak Tersedia")
	}
	dataObject := data.(map[string]interface{})
		return dataObject, nil
}