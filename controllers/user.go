package controllers

import (
	"FlyBlog/models"
	"FlyBlog/utils/logutil"
	"strconv"
)

type UserController struct {
	BaseController
}

//@router /user/:id [get]
func (this *UserController) GoUserHome() {
	suId := this.Ctx.Input.Param(":id")
	uId, err := strconv.Atoi(suId)
	if err != nil {
		logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "GoUserHome", err.Error()})
		this.Ctx.Abort(500, err.Error())
	}
	user := models.GetUser(uId)
	if user == nil {
		logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "GoUserHome", "该用户不存在"})
		this.Ctx.Abort(500, "该用户不存在")
	} else {
		user.Category, _ = models.GetAllCategory(uId)
		this.Data["selectUser"] = user
		this.TplName = "user/home.html"
		return
	}
}

//@router /user/setting/:id [get]
func (this *UserController) Setting() {
	suId := this.Ctx.Input.Param(":id")
	uId, err := strconv.Atoi(suId)
	if err != nil {
		logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "GoUserHome", err.Error()})
		this.Ctx.Abort(500, err.Error())
	}
	user := models.GetUser(uId)
	if user == nil {
		logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "GoUserHome", "该用户不存在"})
		this.Ctx.Abort(500, "该用户不存在")
	} else {
		//user.Category,_=models.GetAllCategory(uId);
		this.Data["selectUser"] = user
		this.TplName = "user/setting.html"
		return
	}
}
