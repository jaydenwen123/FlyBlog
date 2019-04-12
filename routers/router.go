package routers

import (
	"FlyBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "get:Home")
	beego.Router("/about", &controllers.IndexController{}, "get:About")
	beego.Router("/message", &controllers.MessageController{}, "get:Message")
	beego.Include(&controllers.CommentController{})
	beego.Include(&controllers.TestController{})
	beego.Include(&controllers.LoginController{})
	beego.Include(&controllers.ArticleController{})
	beego.Include(&controllers.CategoryController{})
	beego.Include(&controllers.UserController{})
}
