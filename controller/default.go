package controller

import (
	"elearning/model"
	"net/http"
	framework "swikefw"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Default(c *gin.Context) {

	session := sessions.Default(c)
	v := session.Get(framework.Config("jwtName"))
	login := false
	if v != nil {
		login = true
	}

	c.HTML(http.StatusOK, "default/index", gin.H{
		"title":  "Beranda",
		"method": "ok",
		"login":  login,
	})
}

func Login(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get(framework.Config("jwtName"))
	login := false
	if v != nil {
		login = true
	}

	js := []string{
		"/asset/js/popper.min",
		"/asset/js/login",
	}

	css := []string{
		"/asset/css/login",
		"/asset/css/iofrm-style-login",
		"/asset/css/iofrm-theme1-login",
	}

	c.HTML(http.StatusOK, "default/login", gin.H{
		"title": "Login",
		"js":    js,
		"css":   css,
		"login": login,
	})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(framework.Config("jwtName"))
	session.Save()
	c.Redirect(http.StatusFound, "/")
}

func LoginProses(c *gin.Context) {
	data, exist := c.Get("jwt")
	if exist == true || data != nil {
		c.Set("jwt", nil)
	}

	email := c.DefaultPostForm("email", "")
	password := c.DefaultPostForm("password", "")
	user, err := model.Login(email, password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 300,
			"msg":  err.Error(),
		})
		return
	}
	if len(user) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 300,
			"msg":  "Data tidak tersedia",
		})
		return
	}
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := sign.Claims.(jwt.MapClaims)
	claims["id"] = user["id"]
	claims["nama"] = user["nama"]
	claims["username"] = user["username"]
	claims["email"] = user["email"]
	claims["status"] = user["status"]
	claims["tipe"] = user["tipe"]
	token, err := sign.SignedString([]byte(framework.Config("jwtKey")))
	if err != nil {
		framework.ErrorJson("JWT "+err.Error(), c)
		return
	}
	//sesion := sessions.Default(c)
	//sesion.Set(framework.Config("jwtName"), token)
	//sesion.Save()
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "correct login",
		"token": token,
	})
}

