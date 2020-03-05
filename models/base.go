package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type PageInfo struct {
	PageSize int `form:"pageSize"`
	PageNow  int `form:"pageNow"`
	Total    int `form:"total"`
	PageMax  int `form:"pageMax"`
}

const (
	TABLE_ARTICLE="article"
	TABLE_COMMENT="comment"
	TABLE_CATEGORY="category"
	TABLE_USER="user"
	)

var (
	db orm.Ormer
)

//主要注册数据库、注册模型、同时同步数据库，并且最后初始化一个全局操作数据库的对象db
func init() {
	//1.set default database
	err := orm.RegisterDataBase("default", "mysql",
		"root:wen6224261995@tcp(127.0.0.1:3306)/flyblog?charset=utf8", 30)
	if err != nil {
		logs.Error("orm.RegisterDataBase error:", err.Error())
	}
	//2.register model
	orm.RegisterModel(new(User), new(Article), new(Category),new(Comment))

	//3.create table
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Error("orm.RunSyncdb error:", err.Error())
	}
	//4.create interface to finish insert 、delete、update、select
	db = orm.NewOrm()

}
//use mysql;
//update user set host='%' ,password=password('123345')   where user='root';
//
///*
//   ##[mysql 5.7 的修改语句, 5.7 没有 password 字段了，变成了 authentication_string ]*/
//update user set host='%' ,authentication_string=password('123345')  where user='root';
//
//grant  all privileges on *.*  to  root@'%'  identified  by '123456' with  grant  option;
//flush privileges;
