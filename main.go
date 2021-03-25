package main

import (
	_ "ncc/routers"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

