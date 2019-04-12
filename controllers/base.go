package controllers

import (
	"FlyBlog/models"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type UserInfo struct {
	Login bool
	User  models.User
}

type BaseController struct {
	beego.Controller
	UserInfo
}

type HomeArticleVO struct {
	Article      *models.Article
	CanOperation bool
}

func (this *BaseController) Prepare() {
	curUrl := this.Ctx.Input.URI()
	this.Data["path"] = curUrl
	//fmt.Println(this.Ctx.Input.URI())
	this.CheckLogin()
	if !this.Login && (strings.Contains(curUrl, "category") ||
		strings.Contains(curUrl, "article") ||
		strings.Contains(curUrl, "message") ||
		strings.Contains(curUrl, "comment")) {
		this.Redirect("/login", 302)
		return
	}
}

//@router /404 [get,post]
func (this *BaseController) Error404() {
	this.TplName = "error/404.html"
}

//@router /500 [get,post]
func (this *BaseController) Error500() {
	this.TplName = "error/500.html"
}

func (this *BaseController) CheckLogin() {
	if user, ok := this.GetSession("loginUser").(models.User); ok {
		this.User = user
		this.Login = true
		this.Data["user"] = user
		this.Data["isLogin"] = this.Login
	} else {
		this.Login = false
		this.Data["isLogin"] = this.Login
		fmt.Println("the user is not login")
		//this.TplName="login.html"
	}

}
