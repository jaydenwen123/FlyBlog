package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "ToAddArticle",
            Router: `/article`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "AddArticle",
            Router: `/article`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "DeleteArticle",
            Router: `/article/del/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "ShowArticle",
            Router: `/article/show/?:id(\d+)`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "ToUpdateArticle",
            Router: `/article/update/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "UpdateArticle",
            Router: `/article/update/:id`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:BaseController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:BaseController"],
        beego.ControllerComments{
            Method: "Error404",
            Router: `/404`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:BaseController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:BaseController"],
        beego.ControllerComments{
            Method: "Error500",
            Router: `/500`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "ToAddCategory",
            Router: `/category`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "AddCategory",
            Router: `/category`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "DeleteCategory",
            Router: `/category/del/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "ShowCategory",
            Router: `/category/show/?:id(\d+)`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "ToEditCategory",
            Router: `/category/update/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "EditCategory",
            Router: `/category/update/:id`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "ShowArticlesOfCategory",
            Router: `category/:catId/articles`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:CommentController"],
        beego.ControllerComments{
            Method: "ToComment",
            Router: `/comment`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:CommentController"],
        beego.ControllerComments{
            Method: "AddComment",
            Router: `/comment`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:LoginController"],
        beego.ControllerComments{
            Method: "ToLogin",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:LoginController"],
        beego.ControllerComments{
            Method: "LoginOut",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:LoginController"],
        beego.ControllerComments{
            Method: "ToRegister",
            Router: `/register`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:TestController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:TestController"],
        beego.ControllerComments{
            Method: "Test1",
            Router: `/test1`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:TestController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:TestController"],
        beego.ControllerComments{
            Method: "Test2",
            Router: `/test2`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:UserController"],
        beego.ControllerComments{
            Method: "GoUserHome",
            Router: `/user/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/jaydenwen123/FlyBlog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Setting",
            Router: `/user/setting/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
