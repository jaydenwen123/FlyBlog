package controllers

import (
	"FlyBlog/utils/logutil"
	"strconv"
)

type CommentController struct {
	BaseController
}

//@router /comment  [get]
func (this *CommentController) ToComment() {
	//解析文章Id
	artId, err := this.GetInt("artId")

	if err != nil {
		this.Data["addCommentMsg"] = "无法评论文章，文章Id不对"
		this.Redirect("/article/show/"+strconv.Itoa(artId), 302)
		return
	}
	//组装文章和能否操作文章的结构体
	artVO, err := buildArticleVOByArticleId(int64(artId), this.User.Id)
	if err != nil {
		logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "buildArticleVOByArticleId", err.Error()})
		this.Redirect("/", 302)
		return
	}
	this.Data["artVO"] = artVO
	this.TplName = "comment/comment.html"
}

//@router /comment  [post]
func (this *CommentController) AddComment() {
	artId, err := this.GetInt("artId")
	if err != nil {
		this.Data["addCommentMsg"] = "给文章添加评论失败，文章Id不对"
	}
	comUserId, err := this.GetInt("comUserId")
	if err != nil && !this.Login {
		//this.Data["addCommentMsg"]="给文章添加评论失败，您没有登录"
		this.Redirect("/login", 302)
		return
	}
	commentContent := this.GetString("commentContent")
	logutil.Info(logutil.LogInfo{logutil.CurrentFileName(), "AddComment",
		"commUserId:" + strconv.Itoa(comUserId) + "--->artId:" + strconv.Itoa(artId) + "--->commentContent:" + commentContent})

	this.Redirect("/article/show/"+strconv.Itoa(artId), 304)
	//this.TplName = "comment.html"
}
