package controller

import (
	"elearning/helpers"
	"elearning/lib"
	"elearning/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func BuatKelas(c *gin.Context){
	dataJwt, e := lib.GetDataJwt(c)
	if e != nil {
		lib.JsonError(c, e)
		return
	}

	params := model.ParamsKelas{}
	params.Pengajar = dataJwt["id"].(string)
	if err := c.Bind(&params); err != nil{
		lib.JsonError(c, err)
		return
	}

	e = model.BuatKelas(params)
	if e != nil {
		lib.JsonError(c, e)
		return
	}

	lib.Json(c,http.StatusOK, "success", gin.H{})
}

func GetKelasYangDiAjar(c *gin.Context){
	dataJwt, e := lib.GetDataJwt(c)
	if e != nil {
		lib.JsonError(c, e)
		return
	}

	KelasYangDiajar, e := model.GetKelasYangDiajar(dataJwt["id"].(string))
	if e != nil {
		lib.JsonError(c, e)
		return
	}

	lib.Json(c, http.StatusOK, "success", KelasYangDiajar)
}

func UploadMateri(c *gin.Context){
	dataJwt, e := lib.GetDataJwt(c)
	if e != nil {
		lib.JsonError(c, e)
		return
	}
	idkelas := c.PostForm("idkelas")
	judul_materi := c.PostForm("judulmateri")
	filemateri,e := c.FormFile("filemateri")
	if e != nil {
		lib.JsonError(c, e)
		return
	}
	e = helpers.CreateDirIfNotExist("materi", idkelas)
	if e != nil {
		lib.JsonError(c, e)
		return
	}
	filename := dataJwt["id"].(string) +"-"+ filepath.Base(filemateri.Filename)
	if err := c.SaveUploadedFile(filemateri, "public/materi/"+idkelas+"/"+filename); err != nil {
		lib.JsonError(c, err)
		return
	}

	dataToInsert := map[string]interface{}{
		"id_kelas": idkelas,
		"judul_materi": judul_materi,
		"path": "public/materi/"+idkelas+"/"+filename,
	}
	e = model.SaveMateri(dataToInsert)
	if e != nil {
		lib.JsonError(c, e)
		return
	}

	lib.Json(c, http.StatusOK, "success", gin.H{})
}

func GetListMateriThisClass(c *gin.Context){
	idkelas := c.PostForm("idkelas")
	dataMateri, e := model.GetMateriKelasIni(idkelas)
	if e != nil {
		lib.JsonError(c, e)
		return
	}
	lib.Json(c, http.StatusOK, "success", dataMateri)
}

func GetListSiswaThisClass(c *gin.Context){
	idkelas := c.PostForm("idkelas")
	listSiswa, e := model.GetSiswaThisClass(idkelas)
	if e != nil {
		lib.JsonError(c, e)
		return
	}
	lib.Json(c, http.StatusOK, "success", listSiswa)
}
func GetPresensiSiswa(c *gin.Context){
	dataJwt, e := lib.GetDataJwt(c)
	if e != nil {
		lib.JsonError(c, e)
		return
	}
	c.JSON(200,gin.H{"data": dataJwt})
}