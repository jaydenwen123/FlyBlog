package controllers

type IndexController struct {
	BaseController
}

func (c *IndexController) Home() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

func (this *IndexController) About() {
	this.TplName = "about.html"
}
