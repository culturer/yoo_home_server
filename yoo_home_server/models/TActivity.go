package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//家庭活动表
//activityType = false -- FamilyActivity
//activityType = true  -- Homectivity
type TActivity struct {
	Id           int64
	ActivityType bool
	FamilyId     int64
	CreateTime   string
	Desc         string
	AddressId    int64
}

func AddActivity(activityType bool, familyId int64, desc string, addressId int64) (int64, error) {
	activity := &TActivity{ActivityType: activityType, FamilyId: familyId, Desc: desc, CreateTime: time.Now().Format("2006-01-02 15:04:05"), AddressId: addressId}
	o := orm.NewOrm()
	activityId, err := o.Insert(activity)
	return activityId, err
}

func GetActivities(familyId int64, activityType bool) ([]*TActivity, error) {
	activities := make([]*TActivity, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_activity")
	_, err := qs.Filter("family_id", familyId).Filter("ActivityType", activityType).All(&activities)
	return activities, err
}

func DelActivityById(activityId int64) error {
	o := orm.NewOrm()
	activity := &TActivity{Id: activityId}
	_, err := o.Delete(activity)
	return err
}

func UpdateActivity(activityId int64, activityType bool, familyId int64, desc string, addressId int64) error {
	activity := &TActivity{Id: activityId, ActivityType: activityType, FamilyId: familyId, Desc: desc, CreateTime: time.Now().Format("2006-01-02 15:04:05"), AddressId: addressId}
	o := orm.NewOrm()
	_, err := o.Update(activity)
	return err
}
