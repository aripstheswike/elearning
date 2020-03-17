package router

import (
	"elearning/controller"
	"elearning/middleware"
	"github.com/gin-gonic/gin"
)

func RouterUser(r *gin.Engine){
	rUser := r.Group("/siswa")
	rUser.GET("/detail", controller.GetDetailSiswa)
	rUser.Use(middleware.Auth)

	rUser.POST("/join/kelas", controller.JoinKelas)
	rUser.POST("/isi/presensi", controller.MasukKelas)
	rUser.GET("/list/kelas/diikuti", controller.KelasYangDiikuti)
}
