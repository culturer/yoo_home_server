package controllers

import (
	// "github.com/astaxie/beego"
	// "github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
	"yoo_home_server/models"
)

type ActivityController struct {
	BaseController
}

//测试页面
func (this *ActivityController) Get() {
	this.TplName = "activity_test.html"
}

//////////////////////////////////////////////////////////////////////
//																	//
//						获取Activity接口			                //
//																	//
//////////////////////////////////////////////////////////////////////

func (this *ActivityController) Post() {

	options, _ := strconv.ParseInt(this.Input().Get("options"), 10, 64)
	activityType, _ := strconv.ParseBool(this.Input().Get("activityType"))
	familyId, _ := strconv.ParseInt(this.Input().Get("familyId"), 10, 64)
	// [options == 0  查询]
	// [options == 1  增加]
	// [options == 2  删除]
	// [options == 3  修改]
	if options == 0 {

		//activityType = false -- FamilyActivity
		//activityType = true  -- Homectivity
		activities, err := this.getActivities(familyId, activityType)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		this.Data["json"] = map[string]interface{}{"status": 200, "activities": activities, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return

	}

	if options == 1 {
		desc := this.Input().Get("desc")
		addressId, _ := strconv.ParseInt(this.Input().Get("addressId"), 10, 64)
		activityId, err := this.addActivity(activityType, familyId, desc, addressId)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "activityId": activityId, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	if options == 2 {
		activityId, _ := strconv.ParseInt(this.Input().Get("activityId"), 10, 64)
		err := this.delActivity(activityId)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "message": "del activity success", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	if options == 3 {
		activityId, _ := strconv.ParseInt(this.Input().Get("activityId"), 10, 64)
		desc := this.Input().Get("desc")
		addressId, _ := strconv.ParseInt(this.Input().Get("addressId"), 10, 64)
		err := this.updateActivity(activityId, activityType, familyId, desc, addressId)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "message": "update activity success", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	this.Data["json"] = map[string]interface{}{"status": 400, "message": "options error", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

//////////////////////////////////////////////////////////////////////
//																	//
//				将数据库操作再封装一层方便替换model层               //
//																	//
//////////////////////////////////////////////////////////////////////

//获取家庭活动
func (this *ActivityController) getActivities(familyId int64, activityType bool) ([]*models.TActivity, error) {
	activities, err := models.GetActivities(familyId, activityType)
	return activities, err
}

func (this *ActivityController) addActivity(activityType bool, familyId int64, desc string, addressId int64) (int64, error) {
	activityId, err := models.AddActivity(activityType, familyId, desc, addressId)
	return activityId, err
}

func (this *ActivityController) delActivity(activityId int64) error {
	err := models.DelActivityById(activityId)
	return err
}

func (this *ActivityController) updateActivity(activityId int64, activityType bool, familyId int64, desc string, addressId int64) error {
	err := models.UpdateActivity(activityId, activityType, familyId, desc, addressId)
	return err
}
