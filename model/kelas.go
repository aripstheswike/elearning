package model

import (
	"elearning/helpers"
	framework "swikefw"
)

type ParamsKelas struct {
	Id string `form:"id"`
	Pengajar string `form:"pengajar"`
	Namakelas string `form:"namakelas"`
	Kuota string `form:"kuota"`
	Status string `form:"status"`
	Hari string `form:"hari"`
	Deskripsi string `form:"deskripsi"`
}


func BuatKelas(dataKelas ParamsKelas) error {
	dataKelas.Hari = helpers.ConvertHariToDay(dataKelas.Hari)
	db := framework.Database{}
	defer  db.Close()

	db.From("t_kelas")
	_, e :=db.Insert(map[string]interface{}{
		"pengajar": dataKelas.Pengajar,
		"namakelas":dataKelas.Namakelas,
		"hari": dataKelas.Hari,
		"kuota": dataKelas.Kuota,
		"deskripsi": dataKelas.Deskripsi,
	})
	return  e
}

func GetDataUserFlexible(acuan, valueacuan string) (map[string]interface{}, error){
	db := framework.Database{}
	defer  db.Close()

	db.Select("*")
	db.From("t_user")
	db.Where(acuan, valueacuan)
	r,e := db.Row()
	return  r,e
}

func CekStudentRegisteredToClassOrNo(datasiswa map[string]interface{}, idkelas string) (map[string]interface{}, error) {
	db := framework.Database{}
	defer  db.Close()

	db.Select("ku.id")
	db.From("t_kelas_user ku").Join("t_kelas k", "ku.id_kelas=k.id", "")
	db.Where("ku.id_user", datasiswa["id"])
	db.Where("ku.id_kelas", idkelas)
	r,e := db.Row()
	if e != nil{
		return nil, e
	}
	return  r, nil
}


func GetAvailableClass(sortref, order, searchref, searchval string) ([]map[string]interface{}, error) {
	db := framework.Database{}
	defer  db.Close()

	db.Select("k.namakelas, k.deskripsi, k.kuota, k.hari, u.username, DATE_FORMAT(k.date_created, '%D - %M - %Y') date_created").From("t_kelas k").Join("t_user u", "u.id=k.pengajar", "")
	db.Where("k.status", 1).Where("u.tipe", 1)
	if searchref != "" && searchval != ""{
		db.OrderBy(sortref, order)
	}

	if sortref != "" && order != ""{
		db.OrderBy(sortref, order)
	} else {
		db.OrderBy("k.namakelas", "ASC")
	}
	r,e := db.Result()
	return r,e
}

func GetKelasYangDiajar(idpengajar string)([]map[string]interface{}, error){
	db := framework.Database{}
	defer db.Close()

	db.Select("*")
	db.From("t_kelas").Where("pengajar", idpengajar)
	r,e := db.Result()
	return r,e
}

func SaveMateri(datamateri map[string]interface{}) error {
	db := framework.Database{}
	defer db.Close()

	db.From("t_materi")
	_,e := db.Insert(datamateri)
	return e
}

func GetMateriKelasIni(idkelas string)([]map[string]interface{}, error){
	db := framework.Database{}
	defer db.Close()

	db.Select("m.judul_materi, m.path, k.namakelas, k.deskripsi, k.kuota, k.status")
	db.From("t_materi m")
	db.Join("t_kelas k", "m.id_kelas=k.id", "")
	db.Where("m.id_kelas", idkelas).OrderBy("m.date_created", "ASC")

	r,e := db.Result()
	return r,e
}

func GetSiswaThisClass(idkelas string)([]map[string]interface{}, error){
	db := framework.Database{}
	defer db.Close()

	db.Select("ku.id_kelas, ku.id_user, u.username, u.email")
	db.From("t_kelas_user ku")
	db.Join("t_user u", "u.id=ku.id_user", "")
	db.Where("u.tipe", 2)
	db.Where("ku.id_kelas", idkelas).OrderBy("ku.date_created", "ASC")

	r,e := db.Result()
	return r,e
}