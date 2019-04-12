package models

import (
	"FlyBlog/utils/logutil"
	"time"
)

//博客栏目分类
type Category struct {
	Id          int64
	Title       string `orm:"size(256)"`
	Description string `orm:"size(1000)"`
	Visible     bool
	PublishTime time.Time  `orm:"auto_now_add;type(datetime)"`
	UpdateTime  time.Time  `orm:"auto_now;type(datetime)"`
	Articles    []*Article `orm:"reverse(many)"` // 设置一对多的反向关系
	User        *User      `orm:"rel(fk)"`
}

func NewCategory(title string, description string, visible bool, user *User) *Category {
	return &Category{Title: title, Description: description, Visible: visible, User: user}
}

func AddCategory(category *Category) error {
	if _, err := db.Insert(category); err != nil {
		logutil.Error(logutil.LogInfo{FileName: logutil.CurrentFileName(), MethodName: "AddCategory", LogMsg: "failed"})
		return err
	}
	logutil.Debug(logutil.LogInfo{FileName: logutil.CurrentFileName(), MethodName: "AddCategory", LogMsg: "successfully"})
	//logs.Debug("AddCategory sucessfully")
	return nil
}

func UpdateCategory(category *Category) error {
	_, err := db.Update(category)
	return err
}

func DeleteCategory(catid int64) error {
	_, err := db.QueryTable("category").Filter("id", catid).Delete()
	return err
}

func GetCategoryById(catId int64) (*Category, error) {
	category := Category{Id: catId}
	err := db.Read(&category)
	if err != nil {
		return nil, err
	}
	return &category, err
}

func GetAllCategory() ([]*Category, error) {
	var categorys []*Category
	_, err := db.QueryTable("category").OrderBy("-publish_time").All(&categorys)
	return categorys, err
}

func GetCategorysByPage(pageSize, pageNow int, user User) ([]*Category, error) {
	var categorys []*Category
	_, err := db.QueryTable("category").Filter("user_id", user.Id).Limit(pageSize,
		pageSize*(pageNow-1)).OrderBy("-publish_time").All(&categorys)
	return categorys, err
}

func GetCategoryCounts(user User) (int64, error) {
	return db.QueryTable("category").Filter("user_id", user.Id).Count()
}

func GetCategoryWithArticles(catId int64) (*Category, []*Article, error) {
	category := Category{}
	articles := []*Article{}
	err := db.QueryTable("category").Filter("id", catId).RelatedSel().One(&category)
	_, err = db.QueryTable("article").Filter("category_id", catId).RelatedSel().OrderBy("-publish_time").All(&articles)
	return &category, articles, err
}
