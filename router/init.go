package router

import (
	"elearning/controller"
	"net/http"

	"github.com/bandros/framework"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.Static("/asset", "./asset")
	r.Static("/public", "./public")
	r.LoadHTMLGlob("./pages/**/*")
	r.NoRoute(error404)
	r.NoMethod(error404)

	r.GET("/", controller.Dashboard)
	r.GET("/general/get/list/user", controller.GetAllUser)
	r.GET("/general/get/list/kelas", controller.GetListAvailableClass)
	r.POST("/general/register/pengajar", controller.CreatePengajar)
	r.POST("/general/register/siswa", controller.CreateSiswa)
	r.POST("/general/login", controller.LoginProses)

	RouterUser(r)
	RouterPresensi(r)
	RouterJadwal(r)
	RouterMateri(r)

	//r.GET("/", controller.Login)

}

func error404(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get(framework.Config("jwtName"))
	login := false
	if v != nil {
		login = true
	}

	js := []string{
		"/asset/js/popper.min",
		"/asset/js/errorpagenotfound",
	}

	c.HTML(http.StatusNotFound, "error/404", gin.H{
		"title": "Error 404",
		"login": login,
		"js":    js,
	})
}
