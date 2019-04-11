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
	//Author	User	`orm:"rel(fk)"`
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

func GetCategorysByPage(pageSize, pageNow int) ([]*Category, error) {
	var categorys []*Category
	_, err := db.QueryTable("category").Limit(pageSize,
		pageSize*(pageNow-1)).OrderBy("-publish_time").All(&categorys)
	return categorys, err
}

func GetCategoryCounts() (int64, error) {
	return db.QueryTable("category").Count()
}
