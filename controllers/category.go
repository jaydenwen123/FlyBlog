package controllers

import (
	"FlyBlog/models"
	"FlyBlog/utils/logutil"
	"fmt"
	"strconv"
	"strings"
)

type CategoryController struct {
	BaseController
}

//@router /category [get]
func (this *CategoryController) ToAddCategory() {

	this.TplName = "category/category_add.html"
}

//@router /category [post]
func (this *CategoryController) AddCategory() {

	title := this.GetString("title")
	visible := this.GetString("visible")
	description := this.GetString("description")
	isVisible := false
	if len(visible) > 0 && strings.Compare(visible, "on") == 0 {
		isVisible = true
	}
	err := models.AddCategory(models.NewCategory(title,description,isVisible,&this.User))
	//验证栏目信息,检查参数
	if err != nil {
		logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "AddCategory", err.Error()})
		this.Data["addMsg"] = "添加文章栏目失败，请重新添加"
	} else {
		this.Data["addMsg"] = "添加文章栏目成功，请继续添加"
	}
	//添加栏目分类
	//跳转
	this.TplName = "category/category_add.html"
}

//@router /category/show/?:id(\d+) [get]
func (this *CategoryController) ShowCategory() {

	catIds := this.Ctx.Input.Param(":id")
	if catIds != "" && len(strings.TrimSpace(catIds)) > 0 {
		if catId, err := strconv.ParseInt(catIds, 10, 64); err != nil {
			logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "ShowCategory", err.Error()})
			this.Redirect("/", 302)
			return
		} else {
			category, err := models.GetCategoryById(catId)
			if err != nil {
				logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "ShowCategory", err.Error()})
				this.Redirect("/", 302)
				return
			}
			this.Data["category"] = category
			this.TplName = "category/category_list.html"
		}
	} else {
		//按照分页显示数据,获取分页的数据
		pageInfo := models.PageInfo{}
		if err := this.ParseForm(&pageInfo); err != nil {
			fmt.Println("get page info")
			logutil.Error(logutil.LogInfo{logutil.CurrentFileName(),
				"ShowCategory", err.Error()})
			pageInfo.PageNow = 1
			pageInfo.PageSize = 3
		}
		if pageInfo.PageNow == 0 {
			pageInfo.PageNow = 1
		}
		if pageInfo.PageSize == 0 {
			pageInfo.PageSize = 3
		}
		//categories, err := models.GetAllCategory()
		categories, err := models.GetCategorysByPage(pageInfo.PageSize, pageInfo.PageNow)
		total, _ := models.GetCategoryCounts()
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
			fmt.Println(len(categories))

			this.Data["categorys"] = categories
			this.Data["pageMax"] = pageInfo.PageMax
			this.Data["pageNow"] = pageInfo.PageNow
			this.Data["pageSize"] = pageInfo.PageSize
			this.Data["all"] = true
			this.TplName = "category/category_list.html"
		}
	}
}

//@router /category/update/:id [get]
func (this *CategoryController) ToEditCategory() {

	sId := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(sId, 10, 64)
	if err != nil {
		logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "ToEditCategory", err.Error()})
		this.Redirect("/404", 302)
		return
	}
	category, err := models.GetCategoryById(id)
	if err != nil {
		logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "ToEditCategory", err.Error()})
		this.Redirect("/404", 302)
		return
	}
	this.Data["category"] = category
	this.TplName = "category/category_add.html"
}

type CategoryInfo struct {
	Id          int64  `form:"catid"`
	Title       string `form:"title"`
	Visible     string `form:"visible"`
	Description string `form:"description"`
}

//@router /category/update/:id [post]
func (this *CategoryController) EditCategory() {

	//this.Ctx.WriteString("ssss")
	var cateInfo = CategoryInfo{}
	if err := this.ParseForm(&cateInfo); err != nil {
		logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "EditCategory", err.Error()})
		this.Data["updateMsg"] = "更新栏目失败，请重试"
	} else {
		category := models.Category{
			Id:          cateInfo.Id,
			Title:       cateInfo.Title,
			Description: cateInfo.Description,
			Visible:     cateInfo.Visible == "on",
			User:&this.User}
		if err = models.UpdateCategory(&category); err != nil {
			logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "EditCategory", err.Error()})
			this.Data["updateMsg"] = "更新栏目失败，请重试"
		} else {
			this.Data["updateMsg"] = "更新栏目成功"
			this.Data["category"] = category
		}
		//this.Redirect("/category/show",302)
	}
	this.TplName = "category/category_add.html"
}

//返回给前台的处理的信息
type JsonInfo struct {
	Msg    string `json:"msg"`
	Action string `json:"action"`
}

//后期考虑前台全部用ajax来提交请求，统一交由ajax处理
//@router /category/del/:id [get]
func (this *CategoryController) DeleteCategory() {

	//this.Ctx.WriteString("delete category")
	if this.IsAjax() {
		fmt.Println("is ajax")
		//定义一个map用来存放返回给客户端的json数据
		var JsonInfo = JsonInfo{}
		sId := this.Ctx.Input.Param(":id")
		id, err := strconv.ParseInt(sId, 10, 64)
		if err != nil {
			logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "DeleteCategory", err.Error()})
			JsonInfo.Msg = "删除失败"
			JsonInfo.Action = "/404"
		} else {
			if err := models.DeleteCategory(id); err != nil {
				logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "DeleteCategory", err.Error()})
				JsonInfo.Msg = "删除失败"
				JsonInfo.Action = "/404"
			} else {
				JsonInfo.Msg = "删除成功"
				JsonInfo.Action = "/category/show"
			}
		}
		//如果采用json传值，data中的键值只能使用json，不能为其他的键值
		this.Data["json"] = &JsonInfo
		this.ServeJSON()
	}
}

// @router category/:catId/articles [get]
func (this *CategoryController) ShowArticlesOfCategory()  {
	sId := this.Ctx.Input.Param(":catId")
	catId, err := strconv.ParseInt(sId, 10, 64)
	if err != nil {
		logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "ShowArticlesOfCategory", err.Error()})
		this.Ctx.Abort(500,"非法的栏目Id，请重新检查后执行操作")
		return
	}
	category,articles, err := models.GetCategoryWithArticles(catId)
	if err != nil {
		logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "ShowArticlesOfCategory", err.Error()})
		this.Ctx.Abort(500,"没有对应的栏目信息，请重新检查后执行操作")
		return
	}
	this.Data["category"]=category
	this.Data["articles"]=articles
	this.TplName="category/category_article_list.html"
}
