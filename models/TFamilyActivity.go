package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//家庭活动表
type TFamilyActivity struct {
	Id         int64
	FamilyId   int64
	CreateTime string
	Desc       string
	AddressId  int64
}

func AddFamilyActivity(familyActivity *TFamilyActivity) error {
	o := orm.NewOrm()
	_, err := o.Insert(familyActivity)
	return err
}

func GetFamilyActivities(familyId int64) ([]*TFamilyActivity, error) {
	activities := make([]*TFamilyActivity, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_family_activity")
	_, err := qs.Filter("family_id", familyId).All(&activities)
	return activities, err
}
