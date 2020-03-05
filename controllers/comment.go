package controllers

import (
	"github.com/jaydenwen123/FlyBlog/models"
	"github.com/jaydenwen123/FlyBlog/utils/logutil"
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
	jsonInfo:=JsonInfo{}
	if !this.IsAjax(){
		jsonInfo.Msg="请求错误，请稍后重试"
		this.Data["json"]=&jsonInfo
		this.ServeJSON()
		return
	}
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

	//保存评论
	comment:=&models.Comment{}
	comment.Article,_=models.GetArticle(int64(artId))
	comment.Content=commentContent
	comment.CommentUser=models.GetUser(comUserId)
	comment.IsDelete=false

	if err := models.AddComment(comment);err!=nil{
		jsonInfo.Msg="文章评论失败，请重试"
	}else{
		jsonInfo.Msg="文章评论成功，是否返回阅读文章？"
		jsonInfo.Action="/article/show/"+strconv.Itoa(artId)
	}
	//this.Redirect("/article/show/"+strconv.Itoa(artId), 302)
	//this.TplName = "comment/comment.html"
	this.Data["json"]=&jsonInfo
	this.ServeJSON()
}
