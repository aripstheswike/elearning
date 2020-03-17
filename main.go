package main

import (
	"elearning/router"

	framework "github.com/aripstheswike/swikefw"
)

func main() {
	framework := framework.Init{}
	framework.Get()

	r := framework.Begin
	router.Router(r)
	framework.Run()
}
