package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type TActivityItem struct {
	Id               int64
	FamilyId         int64
	HomeActivityId   int64
	FamilyActivityId int64
	Title            string
	Time             string
	Desc             string
}

//查询数据

func GetActivityItemById(id int64) (*TActivityItem, error) {
	o := orm.NewOrm()
	activityItem := new(TActivityItem)
	qs := o.QueryTable("t_activity_item")
	err := qs.Filter("id", id).One(activityItem)
	return activityItem, err
}

func GetActivityItemByFamilyId(familyId int64) ([]*TActivityItem, error) {
	activityItems := make([]*TActivityItem, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_activity_item")
	_, err := qs.Filter("family_id", familyId).All(&activityItems)
	return activityItems, err
}

func GetActivityItemByHomeActivityId(homeActivityId int64) ([]*TActivityItem, error) {
	activityItems := make([]*TActivityItem, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_activity_item")
	_, err := qs.Filter("home_activity_id", homeActivityId).All(&activityItems)
	return activityItems, err
}

func GetActivityItemByFamilyActivityId(familyActivityId int64) ([]*TActivityItem, error) {
	activityItems := make([]*TActivityItem, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_activity_item")
	_, err := qs.Filter("family_activity_id", familyActivityId).All(&activityItems)
	return activityItems, err
}

//增加数据
func AddActivityItem(activityItem *TActivityItem) (int64, error) {
	o := orm.NewOrm()
	activityItemId, err := o.Insert(activityItem)
	return activityItemId, err
}

func DelActicityItem(activityItemId int64) error {
	return nil
}

func UpdateActivityItem(activityItem *TActivityItem) error {
	return nil
}
