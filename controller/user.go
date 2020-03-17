package controller

import (
	"elearning/helpers"
	"elearning/lib"
	"elearning/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func CreateSiswa(c *gin.Context){
	userStruct := model.ParamsUser{}
	if err := c.Bind(&userStruct); err != nil{
		fmt.Println(err.Error())
		lib.JsonError(c, err)
		return
	}
	if model.CekIfUserIsExist(userStruct){
		lib.JsonError(c, errors.New("User Tidak Terdaftar"))
		return
	}
	err := model.TambahDataUser(userStruct, "2")
	if err != nil {
		fmt.Println(err.Error())
		lib.JsonError(c, err)
		return
	}
	lib.Json(c, http.StatusOK, "success", gin.H{})
}

func CreatePengajar(c *gin.Context){
	userStruct := model.ParamsUser{}
	if err := c.Bind(&userStruct); err != nil{
		fmt.Println(err.Error())
		lib.JsonError(c, err)
		return
	}
	if model.CekIfUserIsExist(userStruct){
		lib.JsonError(c, errors.New("User Tidak Terdaftar"))
		return
	}
	err := model.TambahDataUser(userStruct, "1")
	if err != nil {
		fmt.Println(err.Error())
		lib.JsonError(c, err)
		return
	}
	lib.Json(c, http.StatusOK, "success", gin.H{})
}

func GetAllUser(c *gin.Context){
	tipe := c.Query("tipe")
	AllUser, err := model.GetAllUser(tipe)
	if err != nil {
		fmt.Println(err.Error())
		lib.JsonError(c, err)
		return
	}

	lib.Json(c, http.StatusOK,"success", AllUser)
}

func JoinKelas(c *gin.Context){
	emailpelajar := c.PostForm("emailpelajar")
	idkelas := c.PostForm("idkelas")

	datasiswa, e :=  model.GetDataUserFlexible("email", emailpelajar)
	if e != nil {
		fmt.Println("1")
		lib.JsonError(c,e)
		return
	}
	//	cek apakah email sudah terdaftar atau belum
	sudahJoin, e := model.CekStudentRegisteredToClassOrNo(datasiswa,idkelas)
	if e != nil{
		fmt.Println("2")
		lib.JsonError(c,e)
		return
	}
	fmt.Println("sudahJoin ", sudahJoin, " ", len(sudahJoin))
	if len(sudahJoin) == 1{
		fmt.Println("3")
		lib.JsonError(c, errors.New("Anda Sudah Tergabung Di Kelas Ini"))
		return
	}
	//	cek apakah kelas sudah penuh atau belum
	kuota, _ := model.CekSisaKuotaKelasByIdKelas(idkelas)
	if helpers.StrToInt(kuota["kuotasisa"].(string)) == helpers.StrToInt(kuota["kuotaterpenuhi"].(string)){
		fmt.Println("3")
		lib.JsonError(c, errors.New("Kuota Sudah Penuh"))
		return
	}
	//Daftar ke kelas yang dituju
	e = model.JoinKeKelas(datasiswa, idkelas)
	if e != nil{
		fmt.Println("4")
		lib.JsonError(c, e)
		return
	}

	lib.Json(c,http.StatusOK,"success", gin.H{})
}

func GetListAvailableClass(c *gin.Context){
	orderref := c.Query("orderref")
	orderval := c.Query("orderval")
	searchref := c.Query("serchref")
	searchval := c.Query("searchval")
	ListClass, e := model.GetAvailableClass(orderref, orderval, searchref, searchval)
	if e != nil{
		lib.JsonError(c,e)
		return
	}
	lib.Json(c, http.StatusOK, "success", ListClass)
}

func MasukKelas(c *gin.Context){
	idkelas := c.PostForm("idkelas")
	dataSiswa, exist := c.Get("jwt")
	if !exist{
		lib.JsonError(c, errors.New("Anda Belum Login"))
		return
	}
	ObjectDataSiswa := dataSiswa.(map[string]interface{})

//	Cek apakah siswa sudah terdaftar atau belum di kelas tersebut
	joine, e := model.CekStudentRegisteredToClassOrNo(ObjectDataSiswa, idkelas)
	if e != nil{
		lib.JsonError(c, e)
		return
	}

	if len(joine) == 0 {
		lib.JsonError(c, errors.New("Anda belum tergabung di kelas ini!!!"))
		return
	}
//Cek sudah saatnya masuk kelas atau belum
	hariKelasHarusMasuk := joine["hari"].(string)
	hariIni := strings.ToLower(time.Now().Format("Monday"))

	if hariIni != hariKelasHarusMasuk {
		lib.JsonError(c, errors.New("Kelas Belum Dibuka!!!"))
		return
	}

//	isi absesnsi kelas
	e = model.IsiPresensiKelas(ObjectDataSiswa, joine["id"].(string))
	if e != nil{
		lib.JsonError(c, e)
		return
	}

	lib.Json(c,http.StatusOK, "success", gin.H{})
}

func KelasYangDiikuti(c *gin.Context){
	dataJwt, e := lib.GetDataJwt(c)
	fmt.Println("dataJwt =>", dataJwt)
	if e != nil {
		lib.JsonError(c, e)
		return
	}
	followedKelas, e := model.FollowedKelas(dataJwt["id"].(string))
	if e != nil {
		lib.JsonError(c, e)
		return
	}
	lib.Json(c, http.StatusOK, "success", followedKelas)
}

func GetDetailSiswa(c *gin.Context){
	idsiswa := c.Query("idsiswa")
	data, e := model.GetDetailSiswa(idsiswa)
	if e != nil {
		lib.Json(c, http.StatusInternalServerError, e.Error(), gin.H{})
		return
	}
	lib.Json(c, http.StatusOK, "success", data)
}