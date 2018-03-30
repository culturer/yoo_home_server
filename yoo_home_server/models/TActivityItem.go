package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type TActivityItem struct {
	Id         int64
	FamilyId   int64
	ActivityId int64
	Title      string
	CreateTime string
	Desc       string
	Num        int
}

//查询数据

func GetActivityItemById(activityItemId int64) (*TActivityItem, error) {
	o := orm.NewOrm()
	activityItem := new(TActivityItem)
	qs := o.QueryTable("t_activity_item")
	err := qs.Filter("id", activityItemId).One(activityItem)
	return activityItem, err
}

func GetActivityItemByFamilyId(familyId int64) ([]*TActivityItem, error) {
	activityItems := make([]*TActivityItem, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_activity_item")
	_, err := qs.Filter("family_id", familyId).All(&activityItems)
	return activityItems, err
}

func GetActivityItemByActivityId(activityId int64) ([]*TActivityItem, error) {
	activityItems := make([]*TActivityItem, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_activity_item")
	_, err := qs.Filter("activity_id", activityId).All(&activityItems)
	return activityItems, err
}

//增加数据
func AddActivityItem(familyId int64, activityId int64, title string, desc string, createTime string, num int) (int64, error) {
	activityItem := &TActivityItem{FamilyId: familyId, ActivityId: activityId, Title: title, Desc: desc, CreateTime: createTime, Num: num}
	o := orm.NewOrm()
	activityItemId, err := o.Insert(activityItem)
	return activityItemId, err
}

func DelActicityItemById(activityItemId int64) error {
	o := orm.NewOrm()
	activityItem := &TActivityItem{Id: activityItemId}
	_, err := o.Delete(activityItem)
	return err
}

func UpdateActivityItem(activityItemId int64, familyId int64, activityId int64, title string, desc string) error {
	activityItem := &TActivityItem{Id: activityItemId, FamilyId: familyId, ActivityId: activityId, Title: title, CreateTime: time.Now().Format("2006-01-02 15:04:05"), Desc: desc}
	o := orm.NewOrm()
	_, err := o.Update(activityItem)
	return err
}
