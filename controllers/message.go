package controllers

type MessageController struct {
	BaseController
}

func (this *MessageController) Message() {
	this.TplName = "message.html"
}
