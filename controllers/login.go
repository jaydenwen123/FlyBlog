package controllers

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/jaydenwen123/FlyBlog/models"
	"strings"
)

//定义结构体，然后直接通过parseForm方法来解析和传值从前台传递的数据
type UserRegisterInfo struct {
	Username  string `form:"username"`
	Password  string `form:"password"`
	Password2 string `form:"password2"`
	Email     string `form:"email"`
	Role      int    `form:"role"`
}

type LoginController struct {
	BaseController
}

//@router /login [get]
func (this *LoginController) ToLogin() {
	//获取cookie的值
	email := this.Ctx.GetCookie("email")
	password := this.Ctx.GetCookie("password")
	if email != "" {
		this.Data["email"] = email
	}
	if password != "" {
		this.Data["password"] = password
	}
	this.TplName = "login.html"
}

//@router /login [post]
func (this *LoginController) Login() {
	email := this.GetString("email")
	password := this.GetString("password")
	isremember := this.GetString("isremember")
	logs.Info("email:", email, "\tpassword:", password, "\tisremember:", isremember)
	// 处理用户登录请求
	//1.验证用户名和密码
	if user, err := models.CheckUser(email, password); err == nil {
		//2.1将用户信息保存到session中
		logs.Debug("login successfully.", fmt.Sprintf("%#v", user))
		this.SetSession("loginUser", user)
		//2.2如果需要记住密码的话，采用cookie来实现
		if len(isremember) > 0 && strings.Compare(isremember, "on") == 0 {
			this.Ctx.SetCookie("email", user.Email, 100, "/")
			this.Ctx.SetCookie("password", user.Password, 100, "/")
		}
		//2.3.验证成功跳转首页
		this.Redirect("/", 302)
		return
	} else {
		//logs.Error("login failed")
		//3.验证失败跳转404 or 500
		this.Data["loginInfo"]="用户名或者密码不对，请重新检查后输入"
		this.TplName="login.html"
	}

}

//@router /register [get]
func (this *LoginController) ToRegister() {
	this.TplName = "register.html"
}

//@router /register [post]
func (this *LoginController) Register() {
	//1.校验是否用户邮箱已经存在
	regUser := UserRegisterInfo{}
	//解析用户注册的信息
	if !parseRegisterInfo(this, &regUser) {
		return
	}
	if !checkRegisterUserInfo(&regUser, this) {
		return
	}
	//2.然后插入数据库
	if userIsRegistered(&regUser, this) {
		return
	}
	user := models.User{Username: regUser.Username,
		Password: regUser.Password,
		Email:    regUser.Email,
		Role:     regUser.Role}
	if registerFailed(user, this) {
		return
	}
	logs.Info("register user successfully!!!")
	//3.1注册成功重定向到登录页面
	//this.Redirect("/login", 302)
	this.SetSession("regMsg", "注册成功，请登录")
	//this.TplName="login.html"
	this.Redirect("/login", 302)
}

func registerFailed(user models.User, this *LoginController) bool {
	if _, err := models.AddUser(&user); err != nil {
		//3.2注册失败重定向到注册页面
		logs.Error("register user failed", err.Error())
		this.Data["errMsg"] = "注册用户失败，请检查是否有非法输入"
		//this.Redirect("/register",302)
		this.TplName = "register.html"
		return true
	}
	return false
}

func parseRegisterInfo(this *LoginController, regUser *UserRegisterInfo) bool {
	if err := this.ParseForm(regUser); err != nil {
		logs.Error("register info has error", err.Error())
		this.Redirect("/404", 302)
		return false
	}
	return true
}

func userIsRegistered(regUser *UserRegisterInfo, this *LoginController) bool {
	if models.UserIsRegistered(regUser.Email) {
		logs.Error("the email is already registered!!!")
		this.Data["errMsg"] = "用户邮箱已注册"
		//this.Redirect("/register",302)
		this.TplName = "register.html"
		return true
	}
	return false
}

func checkRegisterUserInfo(regUser *UserRegisterInfo, this *LoginController) bool {
	//两次输入的密码不一致
	if strings.Compare(regUser.Password, regUser.Password2) != 0 {
		logs.Error("the password twice is not equal")
		this.Data["errMsg"] = "两次输入的密码不一致"
		//this.Redirect("/register",302)
		this.TplName = "register.html"
		return false
	}
	return true
}

//@router /logout [get]
func (this *LoginController) LoginOut() {
	this.DelSession("loginUser")
	this.Redirect("/", 302)
}
