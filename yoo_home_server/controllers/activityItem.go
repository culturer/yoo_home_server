package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
	"yoo_home_server/models"
)

//活动详情接口
type ActivityItemController struct {
	BaseController
}

//测试页面
func (this *ActivityItemController) Get() {
	this.TplName = "activity_item_test.html"
}

//////////////////////////////////////////////////////////////////////
//																	//
//						获取ActivityItem接口			            //
//																	//
//////////////////////////////////////////////////////////////////////

func (this *ActivityItemController) Post() {

	options, err := strconv.Atoi(this.Input().Get("options"))

	if err != nil {
		beego.Info(err.Error())
	}

	//查询 activityItem
	if options == 0 {

		// activityItemType --- 活动类型
		// [ activityItemType == 0 familyId --- HomeActivity ]
		// [ activityItemType == 1 familyId --- FamilyActivity ]
		// [ activityItemType == 2 activityItemId ]
		// [ activityItemType == 3 HomeActivityId ]
		// [ activityItemType == 4 FamilyActivityId ]
		activityType, _ := strconv.Atoi(this.Input().Get("activityType"))
		familyId, _ := strconv.ParseInt(this.Input().Get("familyId"), 10, 64)

		if activityType == 0 {
			//预留家族活动详情接口

		}

		if activityType == 1 {
			//家庭活动详情接口
			activityitems, err := this.getActivityItemByFamilyId(familyId)
			if err != nil {
				this.Data["json"] = map[string]interface{}{"status": 400, "message": "getHomeActivityItemByActivityId fail", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "activityitems": activityitems, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		if activityType == 2 {
			activityItemId, _ := strconv.ParseInt(this.Input().Get("activityItemId"), 10, 64)
			activityItem, err := this.getActivityItemById(activityItemId)
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "activityitem": activityItem, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		if activityType == 3 {
			homeActivityId, _ := strconv.ParseInt(this.Input().Get("homeActivityId"), 10, 64)
			activityItems, err := this.getActivityItemByHomeActivityId(homeActivityId)
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "activityitems": activityItems, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		if activityType == 4 {
			familyActivityId, _ := strconv.ParseInt(this.Input().Get("familyActivityId"), 10, 64)
			activityItems, err := this.getActivityItemByFamilyActivityId(familyActivityId)
			if err != nil {
				beego.Info(err.Error())
				this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "activityitems": activityItems, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

	}

	//增加activityitem
	if options == 1 {

		familyId, _ := strconv.ParseInt(this.Input().Get("familyId"), 10, 64)
		homeActivityId, _ := strconv.ParseInt(this.Input().Get("homeActivityId"), 10, 64)
		familyActivityId, _ := strconv.ParseInt(this.Input().Get("familyActivityId"), 10, 64)
		desc := this.Input().Get("desc")
		mTime := this.Input().Get("time")
		activityItem := &models.TActivityItem{FamilyId: familyId, HomeActivityId: homeActivityId, FamilyActivityId: familyActivityId, Desc: desc, Time: mTime}
		activityItemId, err := this.addActivityItem(activityItem)
		beego.Info(activityItemId)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "activityItemId": activityItemId, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return

	}

	//删除activityitem
	if options == 2 {

		activityItemId, _ := strconv.ParseInt(this.Input().Get("activityItemId"), 10, 64)
		err := this.delActivityItem(activityItemId)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "message": "删除活动项成功", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return

	}

	//修改activityitem
	if options == 3 {

		familyId, _ := strconv.ParseInt(this.Input().Get("familyId"), 10, 64)
		homeActivityId, _ := strconv.ParseInt(this.Input().Get("homeActivityId"), 10, 64)
		familyActivityId, _ := strconv.ParseInt(this.Input().Get("familyActivityId"), 10, 64)
		desc := this.Input().Get("desc")
		mTime := this.Input().Get("time")
		activityItem := &models.TActivityItem{FamilyId: familyId, HomeActivityId: homeActivityId, FamilyActivityId: familyActivityId, Desc: desc, Time: mTime}
		err := this.updateActvityItem(activityItem)
		if err != nil {
			beego.Info(err)
		}

	}
	this.Data["json"] = map[string]interface{}{"status": 400, "message": "getHomeActivityItemByActivityId fail", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return
}

//////////////////////////////////////////////////////////////////////
//																	//
//				将数据库操作再封装一层方便替换model层               //
//																	//
//////////////////////////////////////////////////////////////////////

func (this *ActivityItemController) getActivityItemByFamilyId(familyId int64) ([]*models.TActivityItem, error) {
	activityItems, err := models.GetActivityItemByFamilyId(familyId)
	return activityItems, err
}

func (this *ActivityItemController) getActivityItemById(id int64) (*models.TActivityItem, error) {
	activityItem, err := models.GetActivityItemById(id)
	return activityItem, err
}

func (this *ActivityItemController) getActivityItemByHomeActivityId(homeActivityId int64) ([]*models.TActivityItem, error) {
	activityItems, err := models.GetActivityItemByHomeActivityId(homeActivityId)
	return activityItems, err
}

func (this *ActivityItemController) getActivityItemByFamilyActivityId(familyActivityId int64) ([]*models.TActivityItem, error) {
	activityItems, err := models.GetActivityItemByFamilyActivityId(familyActivityId)
	return activityItems, err
}

func (this *ActivityItemController) addActivityItem(activityItem *models.TActivityItem) (int64, error) {
	activityItemId, err := models.AddActivityItem(activityItem)
	return activityItemId, err
}

func (this *ActivityItemController) delActivityItem(activityItemId int64) error {
	err := models.DelActicityItem(activityItemId)
	return err
}

func (this *ActivityItemController) updateActvityItem(activityItem *models.TActivityItem) error {
	err := models.UpdateActivityItem(activityItem)
	return err
}
