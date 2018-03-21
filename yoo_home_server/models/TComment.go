package models

import (
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
)

type TComment struct {
	Id         int64
	ArticleId  int64
	CommentId  int64
	UserId     int64
	Desc       string
	CreateTime string
}

func GetCommentsByUserId(userId int64) ([]*TComment, error) {
	comments := make([]*TComment, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_comment")
	_, err := qs.Filter("user_id", userId).All(&comments)
	return comments, err
}

func GetCommentsByArticleId(articleId int64) ([]*TComment, error) {
	comments := make([]*TComment, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_comment")
	_, err := qs.Filter("article_id", articleId).All(&comments)
	return comments, err
}

func AddComment(comment *TComment) (int64, error) {
	o := orm.NewOrm()
	commentId, err := o.Insert(comment)
	return commentId, err
}

func DelComment(comment *TComment) error {
	o := orm.NewOrm()
	_, err := o.Delete(comment)
	return err
}
