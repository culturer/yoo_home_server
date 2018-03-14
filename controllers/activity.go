package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
	"yooplus_indication/models"
)

type ActivityController struct {
	BaseController
}

//测试页面
func (this *ActivityController) Get() {
	this.TplName = "login_test.html"
}

//////////////////////////////////////////////////////////////////////
//																	//
//						获取Activity接口			                //
//																	//
//////////////////////////////////////////////////////////////////////

func (this *ActivityController) Post() {
	activityType, _ := strconv.ParseBool(this.Input().Get("activityType"))
	userId, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
	familyId, _ := strconv.ParseInt(this.Input().Get("familyId"), 10, 64)

	//albumType = false -- FamilyActivity
	//albumType = true  -- Homectivity
	if activityType {
		//获取HomeActivity
		beego.Info("Homectivity")
		activities, err := this.getHomeActivities(familyId)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "getHomeActivities fail", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		this.Data["json"] = map[string]interface{}{"status": 200, "activities": activities, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return

	} else {
		beego.Info("FamilyActivity")

		//获取FamilyActivity
		activities, err := this.getFamilyActivities(userId)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "getFamilyActivities fail", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		this.Data["json"] = map[string]interface{}{"status": 200, "activities": activities, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

}

//////////////////////////////////////////////////////////////////////
//																	//
//				将数据库操作再封装一层方便替换model层               //
//																	//
//////////////////////////////////////////////////////////////////////

//获取家庭活动
func (this *ActivityController) getHomeActivities(familyId int64) ([]*models.TFamilyActivity, error) {
	activities, err := models.GetFamilyActivities(familyId)
	return activities, err
}

//预留接口
//获取家族活动，由于家族关系需要根据每个用户来定义，所以传入的参数为userId
func (this *ActivityController) getFamilyActivities(userId int64) ([]*models.TFamilyActivity, error) {
	return nil, nil
}
