package main

import (
	_ "FlyBlog/routers"
	"github.com/astaxie/beego"
)

func main() {
	//如果需要使用session，则首先需要开启session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
