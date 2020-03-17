package helpers

import (
	"math/rand"
	"os"
	"strconv"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}


func StrToInt(data string) int {
	dataReturn, _ := strconv.Atoi(data)
	return dataReturn
}

func IntToStr(data int) string {
	dataReturn := strconv.Itoa(data)
	return dataReturn
}

func CreateDirIfNotExist(tujuan, namadir string) error {
	if _, err := os.Stat("./public/" + tujuan + "/" + namadir); os.IsNotExist(err) {
		err = os.MkdirAll("./public/"+tujuan+"/"+namadir, 0755)
		if err != nil {
			return err
		}
	}
	return  nil
}

func ConvertHariToDay(hari string) string{
	day := ""
	switch hari {
	case "senin":
		day = "monday"
	case "selasa":
		day = "tuesday"
	case "rabu":
		day = "wednesday"
	case "kamis":
		day = "thursday"
	case "jumat":
		day = "friday"
	case "sabtu":
		day = "saturday"
	case "minggu":
		day = "sunday"
	}
	return day
}
