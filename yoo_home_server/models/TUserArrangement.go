package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//个人日程安排
type TUserArrangement struct {
	Id         int64
	UserId     int64
	Desc       string
	CreateTime string
}

func GetUserArrangementByUserId(userId int64) ([]*TUserArrangement, error) {
	o := orm.NewOrm()
	userArrangement := make([]*TUserArrangement, 0)
	qs := o.QueryTable("t_user_arrangement")
	_, err := qs.Filter("user_id", userId).All(&userArrangement)
	return userArrangement, err
}

func GetUserArrangementById(arrangemenId int64) (*TUserArrangement, error) {
	o := orm.NewOrm()
	arrangement := new(TUserArrangement)
	qs := o.QueryTable("t_user_arrangement")
	err := qs.Filter("id", arrangemenId).One(arrangement)
	return arrangement, err

}

func AddUserArrangement(userArrangement *TUserArrangement) (int64, error) {
	o := orm.NewOrm()
	userArrangementId, err := o.Insert(userArrangement)
	return userArrangementId, err
}

func DelUserArrangement(userArrangement *TUserArrangement) error {
	o := orm.NewOrm()
	_, err := o.Delete(userArrangement)
	return err
}

func UpdateUserArrangement(userArrangement *TUserArrangement) error {
	o := orm.NewOrm()
	_, err := o.Update(userArrangement)
	return err
}
