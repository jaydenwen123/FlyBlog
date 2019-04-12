package models

import "time"

type Comment struct {
	Id	int64
	Article	*Article `orm:"rel(fk)"`
	Content string `orm:"size(1000)"`
	LikeCount	int64
	CommentUser	*User `orm:"rel(fk)"`
	PublishTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime  time.Time `orm:"auto_now;type(datetime)"`
	IsDelete	bool	//评论是否删除了
}

func AddComment(comment *Comment)  error{
	_, err := db.Insert(comment)
	return err
}

func UpdateComment(comment *Comment) error {
	_, err := db.Update(comment)
	return err
}

func DeleteComment(commentId int64) error {
	_, err := db.QueryTable(TABLE_COMMENT).Filter("id", commentId).Delete()
	return err
}

//通过分页获取文章artId的评论内容
func GetComments(artId int64) ([] *Comment,error)  {
	comments :=[]*Comment{}
	_, err := db.QueryTable(TABLE_COMMENT).Filter("article_id", artId).
		OrderBy("-like_count").Limit(10,1).
		RelatedSel().All(&comments)
	return comments,err
}

