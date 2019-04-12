package controllers

import (
	"FlyBlog/models"
	"FlyBlog/utils/logutil"
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

type ArticleController struct {
	BaseController
}

type ArticleInfo struct {
	Id       int64  `form:"artId";json:"artId"`
	Title    string `form:"title";json:"title"`
	Category int64  `form:"category";json:"category"`
	Content  string `form:"content";json:"content"`
}

//@router /article [get]
func (this *ArticleController) ToAddArticle() {
	//获取所有的文章栏目，数据初始化文章栏目
	categories, err := models.GetAllCategory(this.User.Id)
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
	//fmt.Println(this.IsAjax())

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
		art.User = &this.User
		art.Summary, _ = getSummary(article.Content)
		logutil.Info(logutil.LogInfo{logutil.CurrentFileName(), "AddArticle", "文章摘要：\n" + art.Summary})
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
	if this.IsAjax(){
		jsonInfo.Action="/article/show"
		this.Data["json"]=&jsonInfo
		this.ServeJSON()
	}else{
	this.Data["addMsg"] = jsonInfo.Msg
	this.ToAddArticle()
	}
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
	if this.IsAjax() {
		logutil.Info(logutil.LogInfo{logutil.CurrentFileName(),
			"UpdateArticle", "使用的是ajax请求提交数据"})
		//this.Data["updateMsg"] = "数据更新失败，请检查数据是否有问题，然后重试"
		jsonInfo := JsonInfo{"数据更新成功,3秒后跳转到文章列表页面", "/article/show"}
		art := ArticleInfo{}
		if err := this.ParseForm(&art); err != nil {
			logutil.Error(logutil.LogInfo{logutil.CurrentFileName(),
				"UpdateArticle", err.Error()})
			jsonInfo.Msg = "数据非法输入，请检查后重新更新"
		} else {
			artCategory, _ := models.GetCategoryById(art.Category)
			summary, _ := getSummary(art.Content)
			logutil.Info(logutil.LogInfo{logutil.CurrentFileName(), "UpdateArticle", "文章摘要：\n" + summary})
			article := models.NewArticle(art.Id, art.Title, art.Content,
				artCategory, &this.User, summary)
			if err := models.UpdateArticle(article); err != nil {
				logutil.Error(logutil.LogInfo{logutil.CurrentFileName(),
					"UpdateArticle", err.Error()})
				jsonInfo.Msg = "数据更新失败，请稍后重试"
			} else {
				jsonInfo.Msg = "数据更新成功,请继续选择其他操作"
			}
			//this.Data["article"] = article
		}
		this.Data["json"] = jsonInfo
		this.ServeJSON()
	} else {

		handleCommonUpdate(this)
	}
}

func handleCommonUpdate(this *ArticleController) {
	art := ArticleInfo{}
	if err := this.ParseForm(&art); err != nil {
		logutil.Error(logutil.LogInfo{logutil.CurrentFileName(),
			"UpdateArticle", err.Error()})
		//this.Data["updateMsg"]="保存文章失败，请重试"
		//this.Redirect("/500", 302)
		this.Data["updateMsg"] = "数据更新失败，请检查数据是否有问题，然后重试"
	} else {
		artCategory, _ := models.GetCategoryById(art.Category)
		//截取文章摘要信息
		summary, _ := getSummary(art.Content)

		logutil.Info(logutil.LogInfo{logutil.CurrentFileName(), "handleCommonUpdate", "文章摘要：\n" + summary})
		article := models.NewArticle(art.Id, art.Title, art.Content,
			artCategory, &this.User, summary)
		if err := models.UpdateArticle(article); err != nil {
			logutil.Error(logutil.LogInfo{logutil.CurrentFileName(),
				"UpdateArticle", err.Error()})
			this.Data["updateMsg"] = "数据更新失败，请稍后重试"
		} else {
			this.Data["updateMsg"] = "数据更新成功,请继续选择其他操作"
		}
		this.Data["article"] = article
	}
	//this.ToUpdateArticle()
	this.TplName = "article/article_pub.html"
}

//@router /article/del/:id [get]
func (this *ArticleController) DeleteArticle() {
	//this.Ctx.WriteString("delete category")
	if this.IsAjax() {
		//fmt.Println("is ajax")
		//定义一个map用来存放返回给客户端的json数据
		var JsonInfo = JsonInfo{}
		aId := this.Ctx.Input.Param(":id")
		id, err := strconv.ParseInt(aId, 10, 64)
		if err != nil {
			logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "DeleteArticle", err.Error()})
			JsonInfo.Msg = "删除失败"
			JsonInfo.Action = "/404"
		} else {
			if err := models.DeleteArticle(id); err != nil {
				logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "DeleteArticle", err.Error()})
				JsonInfo.Msg = "删除失败"
				JsonInfo.Action = "/404"
			} else {
				JsonInfo.Msg = "删除成功"
				JsonInfo.Action = "/article/show"
			}
		}
		//如果采用json传值，data中的键值只能使用json，不能为其他的键值
		this.Data["json"] = &JsonInfo
		this.ServeJSON()
	}
}

func buildArticleVOByArticleId(artId int64, userId int) (*HomeArticleVO, error) {
	article, err := models.GetArticle(artId)
	if err != nil {
		return nil, err
	}
	artVO := HomeArticleVO{}
	artVO.Article = article
	artVO.CanOperation = canOperation(article.User.Id, userId)
	return &artVO, nil
}

//@router /article/show/?:id(\d+) [get]
func (this *ArticleController) ShowArticle() {
	artIds := this.Ctx.Input.Param(":id")
	if artIds != "" && len(strings.TrimSpace(artIds)) > 0 {
		this.ShowOneArticle(artIds)
	} else {
		//按照分页显示数据,获取分页的数据
		pageInfo := models.PageInfo{}
		if err := this.ParseForm(&pageInfo); err != nil {
			fmt.Println("get page info")
			logutil.Error(logutil.LogInfo{logutil.CurrentFileName(),
				"ShowArticle", err.Error()})
			pageInfo.PageNow = 1
			pageInfo.PageSize = 3
		}
		if pageInfo.PageNow == 0 {
			pageInfo.PageNow = 1
		}
		if pageInfo.PageSize == 0 {
			pageInfo.PageSize = 3
		}
		articles, err := models.GetArticlesWithPage(pageInfo.PageSize, pageInfo.PageNow, &this.User)
		total, _ := models.GetArticleCounts(&this.User)
		pageInfo.Total = int(total)
		pageInfo.PageMax = pageInfo.Total / pageInfo.PageSize
		if pageInfo.Total%pageInfo.PageSize != 0 {
			pageInfo.PageMax++
		}

		if err != nil {
			logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "ShowCategory", err.Error()})
			this.Redirect("/", 302)
			return
		} else {
			fmt.Println(pageInfo)
			fmt.Println(len(articles))

			this.Data["articles"] = articles
			this.Data["pageMax"] = pageInfo.PageMax
			this.Data["pageNow"] = pageInfo.PageNow
			this.Data["pageSize"] = pageInfo.PageSize
			this.Data["all"] = true
			this.TplName = "article/article_list.html"
		}
	}
}
//展现一篇文章，同时取出文章的评论
func (this *ArticleController) ShowOneArticle(artIds string) {
	if artId, err := strconv.ParseInt(artIds, 10, 64); err != nil {
		logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "ShowArticle", err.Error()})
		this.Redirect("/", 302)
	} else {
		//组装文章和用户对文章权限的数据
		artVO, err := buildArticleVOByArticleId(artId, this.User.Id)
		if err != nil {
			logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "buildArticleVOByArticleId", err.Error()})
			this.Redirect("/", 302)
		}
		//按照点赞数取评论信息
		comments,err:=models.GetComments(artId)
		if err != nil {
			logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "ShowOneArticle", err.Error()})
			this.Redirect("/", 302)
		}
		this.Data["artVO"] = artVO
		this.Data["comments"]=comments
		this.TplName = "article/article_detail.html"
	}
}

//截取文章的摘要信息
func getSummary(content string) (string, error) {
	var buf bytes.Buffer
	buf.Write([]byte(content))
	doc, err := goquery.NewDocumentFromReader(&buf)
	if err != nil {
		return "", err
	}
	str := doc.Find("body").Text()
	strRune := []rune(str)
	if len(strRune) > 400 {
		strRune = strRune[:400]
	}
	return string(strRune) + "...", nil
}
