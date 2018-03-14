package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type TPhoto struct {
	Id         int64
	AlbumId    int64
	UserId     int64
	FamilyId   int64
	ArticleId  int64
	FileName   string
	FileUrl    string
	CreateTime string
}

func GetPhotos(userId int64) ([]*TPhoto, error) {
	photos := make([]*TPhoto, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_photo")
	_, err := qs.Filter("user_id", userId).All(&photos)
	return photos, err
}

func GetPhotosByArticleId(articleId int64) ([]*TPhoto, error) {
	photos := make([]*TPhoto, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_photo")
	_, err := qs.Filter("article_id", articleId).All(&photos)
	return photos, err
}

func AddPhoto(photo *TPhoto) (int64, error) {
	o := orm.NewOrm()
	photoId, err := o.Insert(photo)
	return photoId, err
}

func DelPhoto(photo *TPhoto) error {
	o := orm.NewOrm()
	_, err := o.Delete(photo)
	return err
}
