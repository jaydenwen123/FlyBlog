package controllers

import (
	"github.com/jaydenwen123/FlyBlog/models"
	"github.com/jaydenwen123/FlyBlog/utils/logutil"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Home() {
	//从数据库按照文章的优先级等数据或者发表时间逆序取数据
	articles, err := models.GetArticlesWithPage(10, 1, nil)
	if err != nil {
		logutil.Error(logutil.LogInfo{logutil.CurrentFileName(), "Home", err.Error()})
		c.Ctx.Abort(500, "服务器内部错误，无法获取数据")
	}
	homeArticleVOs := make([]*HomeArticleVO, 0)
	canOper := false
	if c.Login {
		for index, _ := range articles {
			canOper = canOperation(articles[index].User.Id, c.User.Id)
			homeArticleVOs = append(homeArticleVOs, &HomeArticleVO{articles[index], canOper})
			//fmt.Sprintf("%#v",homeArticleVOs)

		}
	} else {
		for index, _ := range articles {
			//fmt.Println(articles[index])
			homeArticleVOs = append(homeArticleVOs, &HomeArticleVO{articles[index], false})
		}
	}
	//fmt.Println(homeArticleVOs)
	//必须传递的参数
	c.Data["homeArticleVOs"] = homeArticleVOs
	c.TplName = "index.html"
}

func canOperation(artUserId, loginUserId int) bool {
	return artUserId == loginUserId
}

func (this *IndexController) About() {
	this.TplName = "about.html"
}
