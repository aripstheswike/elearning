package middleware

import (
	"elearning/lib"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	framework "swikefw"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	c.Set("jwt", "")
	c.Save()
	var tokenString string
	tokenHeader := c.GetHeader("token")
	if tokenHeader == "" {
		lib.JsonError(c,errors.New("Error Token"))
		c.Abort()
		return
	} else {
		tokenString = tokenHeader
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(framework.Config("jwtKey")), nil
	})
	if token == nil || err != nil {
		fmt.Println("Token Error =>", err.Error())
		lib.Json(c, http.StatusInternalServerError, "forbidden", gin.H{})
		c.Abort()
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		fmt.Println("Token is invalid =>", err.Error())
		lib.Json(c, http.StatusInternalServerError, "forbidden", gin.H{})
		return
	} else {
		data_jwt := map[string]interface{}{}
		if claims["tipe"] != "2"{
			lib.Json(c, http.StatusInternalServerError, "Anda Bukan Pelajar", gin.H{})
			c.Abort()
		}
		data_jwt["jti"] = claims["jti"]
		data_jwt["id"] = claims["id"]
		data_jwt["email"] = claims["email"]
		data_jwt["tipe"] = claims["tipe"]
		data_jwt["username"] = claims["username"]
		c.Set("jwt", data_jwt)
		c.Save()
	}

}


func AuthPengajar(c *gin.Context) {
	c.Set("jwt", "")
	c.Save()
	var tokenString string
	tokenHeader := c.GetHeader("token")
	if tokenHeader == "" {
		lib.JsonError(c,errors.New("Error Token"))
		c.Abort()
		return
	} else {
		tokenString = tokenHeader
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(framework.Config("jwtKey")), nil
	})
	if token == nil || err != nil {
		fmt.Println("Token Error =>", err.Error())
		lib.Json(c, http.StatusInternalServerError, "forbidden", gin.H{})
		c.Abort()
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		fmt.Println("Token is invalid =>", err.Error())
		lib.Json(c, http.StatusInternalServerError, "forbidden", gin.H{})
		return
	} else {
		if claims["tipe"] != "1"{
			lib.Json(c, http.StatusInternalServerError, "Anda Bukan Pengajar", gin.H{})
			c.Abort()
		}
		data_jwt := map[string]interface{}{}
		data_jwt["jti"] = claims["jti"]
		data_jwt["id"] = claims["id"]
		data_jwt["email"] = claims["email"]
		data_jwt["level"] = claims["level"]
		data_jwt["username"] = claims["username"]
		data_jwt["tipe"] = claims["tipe"]
		c.Set("jwt", data_jwt)
		c.Save()
	}


}


func notfound(session sessions.Session, c *gin.Context) {
	// session.Delete(framework.Config("jwtName"))
	// session.Save()
	// c.Redirect(http.StatusFound, "/")
}
