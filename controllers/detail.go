package controllers

type DetailController struct {
	BaseController
}

func (this *DetailController) Detail() {
	this.TplName = "details.html"
}
