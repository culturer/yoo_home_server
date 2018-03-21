package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type TAlbumItem struct {
	Id         int64
	UserId     int64
	FamilyId   int64
	Name       string
	Icon       string
	CreateTime string
}

func GetUserAlbumItem(userId int64) ([]*TAlbumItem, error) {
	albums := make([]*TAlbumItem, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_album_item")
	_, err := qs.Filter("user_id", userId).All(&albums)
	return albums, err
}

func GetFamilyAlbumItem(familyId int64) ([]*TAlbumItem, error) {
	albums := make([]*TAlbumItem, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_album_item")
	_, err := qs.Filter("family_id", familyId).All(&albums)
	return albums, err
}

func GetAlbumItemById(albumItemId int64) (*TAlbumItem, error) {
	o := orm.NewOrm()
	albumItem := new(TAlbumItem)
	qs := o.QueryTable("t_album_item")
	err := qs.Filter("id", albumItemId).One(albumItem)
	return albumItem, err
}

func AddAlbumItem(albumItem *TAlbumItem) (int64, error) {
	o := orm.NewOrm()
	albumItemId, err := o.Insert(albumItem)
	return albumItemId, err
}
func DelAlbumItem(albumItem *TAlbumItem) error {
	o := orm.NewOrm()
	_, err := o.Delete(albumItem)
	return err
}
func UpdateAlbumItem(albumItem *TAlbumItem) error {
	o := orm.NewOrm()
	_, err := o.Update(albumItem)
	return err
}
