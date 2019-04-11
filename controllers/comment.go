package controllers

type CommentController struct {
	BaseController
}

func (this *CommentController) Comment() {
	this.TplName = "comment.html"
}
