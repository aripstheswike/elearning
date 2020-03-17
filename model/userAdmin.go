package model

import (
	"fmt"
	framework "swikefw"
)

func Login(email string, password string) (map[string]interface{}, error) {
	db := framework.Database{}
	defer db.Close()
	db.Select("*")
	db.From("t_user")
	db.Where("email", email)
	db.Where("password", framework.Password(password))
	data, err := db.Row()
	fmt.Println("data login =>", data)
	return data, err
}
