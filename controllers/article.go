package controllers

import (
	"FlyBlog/models"
	"FlyBlog/utils/logutil"
	"fmt"
	"strconv"
)

type ArticleController struct {
	BaseController
}

type ArticleInfo struct {
	Title    string `form:"title"`
	Category int64  `form:"category"`
	Content  string `form:"content"`
}

//@router /article [get]
func (this *ArticleController) ToAddArticle() {
	//获取所有的文章栏目，数据初始化文章栏目
	categories, err := models.GetAllCategory()
	if err != nil {
		logutil.Error(logutil.LogInfo{logutil.CurrentFileName(),
			"ToAddArticle", err.Error()})
		this.Redirect("/login", 302)
	}
	this.Data["categories"] = categories
	this.TplName = "article/article_pub.html"
}

//@router /article [post]
func (this *ArticleController) AddArticle() {
	//采用ajkx提交数据
	article := ArticleInfo{}
	fmt.Println(this.IsAjax())
	//if !this.IsAjax() {
	//	this.Abort("500")
	//}
	jsonInfo := JsonInfo{}
	if err := this.ParseForm(&article); err != nil {
		logutil.Error(logutil.LogInfo{logutil.CurrentFileName(),
			"ToAddArticle", err.Error()})
		jsonInfo.Msg = "文章内容有错误，请检查后重新提交"
	} else {
		//保存文章
		art := models.Article{}
		art.Title = article.Title
		art.Category, _ = models.GetCategoryById(article.Category)
		art.Content = article.Content
		if err := models.AddArticle(&art); err != nil {
			logutil.Error(logutil.LogInfo{logutil.CurrentFileName(),
				"ToAddArticle", err.Error()})
			jsonInfo.Msg = "保存文章失败，请重试"
		} else {
			logutil.Info(logutil.LogInfo{logutil.CurrentFileName(),
				"ToAddArticle", "保存文章成功"})
			jsonInfo.Msg = "保存文章成功，请继续添加文章"
		}
	}
	//jsonInfo.Action="/"
	//this.Data["json"]=&jsonInfo
	//this.ServeJSON()
	this.Data["addMsg"] = jsonInfo.Msg
	this.ToAddArticle()
}

//@router /article/update/:id [get]
func (this *ArticleController) ToUpdateArticle() {
	sId := this.Ctx.Input.Param(":id")
	aId, err := strconv.ParseInt(sId, 10, 64)
	if err != nil {
		logutil.Error(logutil.LogInfo{logutil.CurrentFileName(),
			"ToUpdateArticle", err.Error()})
		//this.Data["updateMsg"]="保存文章失败，请重试"
		this.Redirect("/500", 302)
	}
	if article, err := models.GetArticle(aId); err != nil {
		logutil.Error(logutil.LogInfo{logutil.CurrentFileName(),
			"ToUpdateArticle", err.Error()})
		//this.Data["updateMsg"]="保存文章失败，请重试"
		this.Redirect("/500", 302)
	} else {
		this.Data["article"] = article
		this.Data["add"] = false
	}
	this.ToAddArticle()
	//this.TplName="article/article_pub.html"
}

//@router /article/update/:id [post]
func (this *ArticleController) UpdateArticle() {

	this.TplName = "index.html"
}

//@router /article/del/:id [get]
func (this *ArticleController) DeleteArticle() {
	this.TplName = "index.html"
}

//@router /article/show/?:id(\d+) [get]
func (this *ArticleController) ShowArticle() {
	article, err := models.GetArticlesWithPage(10, 1)
	if err != nil {
		logutil.Error(logutil.LogInfo{logutil.CurrentFileName(),
			"ShowArticle", err.Error()})
		//this.Data["updateMsg"]="保存文章失败，请重试"
		this.Redirect("/500", 302)
	}
	this.Data["articles"] = article
	this.Data["all"] = true
	this.TplName = "article/article_list.html"
}
