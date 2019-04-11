package controllers

type TestController struct {
	BaseController
}

//@router /test1 [get]
func (c *TestController) Test1() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

//@router /test2 [get,post]
func (this *TestController) Test2() {
	this.TplName = "about.html"
}
