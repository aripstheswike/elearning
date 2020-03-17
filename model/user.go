package model

import (
	"fmt"
	framework "swikefw"
)

type   ParamsUser struct{
	Username string `form:"username"`
	Email string `form:"email"`
	Password string `form:"password"`
	Tipe string `form:"tipe"`
	Status string `form:"string"`
}

func TambahDataUser(datauser ParamsUser, tipe string) error {
	db := framework.Database{}
	defer db.Close()
	db.From("t_user")
	_, err:=	db.Insert(map[string]interface{}{
		"username": datauser.Username,
		"email": datauser.Email,
		"password": framework.Password(datauser.Password),
		"tipe": tipe,
	})
	return err
}

func CekIfUserIsExist(datauser ParamsUser) bool {
	db := framework.Database{}
	defer db.Close()
	db.Select("id")
	db.From("t_user")
	db.Where("username LIKE", "%"+ datauser.Username +"%")
	db.Where("email LIKE", "%"+ datauser.Email +"%")
	r, _ := db.Row()
	if len(r) == 0 {
		return false
	}
	return true
}

func GetAllUser(tipe string) ([]map[string]interface{}, error){

	db := framework.Database{}
	defer db.Close()

	db.Select("*")
	db.From("t_user").OrderBy("username", "DESC")
	if tipe != ""{
		db.Where("tipe", tipe)
	}
	r, e := db.Result()
	return r,e
}


func CekSisaKuotaKelasByIdKelas(idkelas string)(map[string]interface{}, error){
	db := framework.Database{}
	defer db.Close()

	db.Select("COUNT(ku.id_user) kuotaterpenuhi, k.kuota kuotasisa").From("t_kelas_user ku")
	db.Join("t_kelas k", "k.id=ku.id_kelas", "").Where("k.id",idkelas)
	data, e := db.Row()
	if e != nil{
		return nil, e
	}
	return data, nil
}

func JoinKeKelas(datasiswa map[string]interface{}, idkelas string) error {
	db := framework.Database{}
	defer db.Close()

	db.From("t_kelas_user")
	_,e :=db.Insert(map[string]interface{}{
		"id_user": datasiswa["id"],
		"id_kelas": idkelas,
	})
	return e
}

func IsiPresensiKelas(datasiswa map[string]interface{}, idkelasuser string) error{
	db := framework.Database{}
	defer db.Close()

	db.From("t_mutasi_absensi")
	_, e:= db.Insert(map[string]interface{}{
		"id_user": datasiswa["id"],
		"id_kelas_user": idkelasuser,
	})
	return e
}
func FollowedKelas(iduser string) ([]map[string]interface{}, error){
	db := framework.Database{}
	defer db.Close()

	db.Select("k.namakelas, u.username pengajar, k.kuota, k.deskripsi")
	db.From("t_kelas_user ku").Join("t_kelas k", "k.id=ku.id_kelas", "")
	db.Join("t_user u", "u.id=k.pengajar", "")
	db.Where("id_user", iduser)
	r,e := db.Result()
	fmt.Println(db.QueryView())
	return r,e
}
func GetDetailSiswa(idsiswa string)([]map[string]interface{}, error){
	db := framework.Database{}
	defer db.Close()

	db.Select("u.username, u.email, u.tipe, u.status, k.namakelas ").From("t_user u")
	db.Join("t_kelas_user ku", "u.id=ku.id_user", "")
	db.Join("t_kelas k", "ku.id_kelas=k.id","")
	db.Where("u.id", idsiswa)
	r,e := db.Result()
	return r,e
}