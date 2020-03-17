package router

import (
	"elearning/controller"
	"elearning/middleware"
	"github.com/gin-gonic/gin"
	)

func RouterJadwal(r *gin.Engine) {
	rKelas := r.Group("/class")
	rKelas.Use(middleware.AuthPengajar)

	rKelas.POST("/create/", controller.BuatKelas)
	rKelas.GET("/my/", controller.GetKelasYangDiAjar)
	rKelas.POST("/upload/materi", controller.UploadMateri)
	rKelas.POST("/get/all/materi", controller.GetListMateriThisClass)
	rKelas.POST("/get/list/siswa", controller.GetListSiswaThisClass)
	rKelas.POST("/get/presensi/siswa", controller.GetPresensiSiswa)
}
