package models

import (
	"time"
)

//博客文章
type Article struct {
	Id          int64
	Title       string    `orm:"size(256)"`
	Content     string    `orm:"size(5000)"`
	PublishTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime  time.Time `orm:"auto_now;type(datetime)"`
	Category    *Category `orm:"rel(fk)" description:(这个是文章的栏目分类)` //设置一对多关系
}

const TABLE_ARTICLE = "article"

func init() {

}

func AddArticle(article *Article) error {
	_, err := db.Insert(article)
	return err
}

func UpdateArticle(article *Article) error {
	_, err := db.Update(article)
	return err
}

func DeleteArticle(aId int64) error {
	_, err := db.QueryTable(TABLE_ARTICLE).Filter("id", aId).Delete()
	return err
}

func GetArticle(aId int64) (*Article, error) {
	article := Article{}
	err := db.QueryTable(TABLE_ARTICLE).Filter("id", aId).One(&article)
	db.LoadRelated(&article, "Category")
	return &article, err
}

func GetArticlesByCategory(cat_id int64) ([]*Article, error) {
	var articles []*Article
	_, err := db.QueryTable(TABLE_ARTICLE).Filter("category_id", cat_id).RelatedSel().All(&articles)
	//for k,_:=range articles{
	//	db.LoadRelated(articles[k],"Category")
	//}
	return articles, err
}

//一对多和多对一、以及一对一关联时，需要使用RelatedSet方法来衔接
func GetArticlesWithPage(pageSize, pageNow int) ([]*Article, error) {
	var articles []*Article
	_, err := db.QueryTable(TABLE_ARTICLE).Limit(pageSize, (pageNow - 1)).RelatedSel().All(&articles)
	return articles, err
}
