package models

import (
	"time"
)

//博客文章
type Article struct {
	Id          int64
	Title       string    `orm:"size(256)"`
	Content     string    `orm:"type(text)"`
	Summary     string    `orm:"type(text)"`
	PublishTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime  time.Time `orm:"auto_now;type(datetime)"`
	Category    *Category `orm:"rel(fk);description:(这个是文章的栏目分类)" ` //设置一对多关系
	User        *User     `orm:"rel(fk)"`
}

func NewArticle(id int64, title string, content string, category *Category, user *User, summary string) *Article {
	return &Article{Id: id, Title: title, Content: content, Category: category, User: user, Summary: summary}
}
func NewArticle2(title string, content string, category *Category, user *User) *Article {
	return &Article{Title: title, Content: content, Category: category, User: user}
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
func GetArticlesWithPage(pageSize, pageNow int, user *User) ([]*Article, error) {
	var articles []*Article
	var err error
	if user != nil {
		_, err = db.QueryTable(TABLE_ARTICLE).Filter("user_id", user.Id).Limit(pageSize, (pageNow - 1)).RelatedSel().OrderBy("-update_time").All(&articles)
	} else {
		_, err = db.QueryTable(TABLE_ARTICLE).Limit(pageSize, (pageNow - 1)).RelatedSel().OrderBy("-update_time").All(&articles)
	}
	return articles, err
}

func GetArticleCounts(user *User) (int64, error) {
	if user != nil {
		return db.QueryTable("article").Filter("user_id", user.Id).Count()
	} else {
		return db.QueryTable("article").Count()

	}
}
