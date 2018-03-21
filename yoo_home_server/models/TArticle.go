package models

import (
	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
)

//亲友圈文章
type TArticle struct {
	Id         int64
	UserId     int64
	AddressId  int64
	CreateTime string
	Dcsc       string
}

func GetArticlesByUserId(userId int64) ([]*TArticle, error) {
	o := orm.NewOrm()
	articles := make([]*TArticle, 0)
	qs := o.QueryTable("t_article")
	_, err := qs.Filter("user_id", userId).All(&articles)
	return articles, err
}

func GetArticleById(articleId int64) (*TArticle, error) {
	o := orm.NewOrm()
	article := new(TArticle)
	qs := o.QueryTable("t_article")
	err := qs.Filter("id", articleId).One(article)
	return article, err
}

func AddArticle(article *TArticle) (int64, error) {
	o := orm.NewOrm()
	articleId, err := o.Insert(article)
	return articleId, err
}

func DelArticle(article *TArticle) error {
	o := orm.NewOrm()
	_, err := o.Delete(article)
	return err
}
