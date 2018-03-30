package controllers

import (
	// "github.com/astaxie/beego"
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
	// [options == 0  查询]
	// [options == 1  增加]
	// [options == 2  删除]
	// [options == 3  修改]
	options, _ := strconv.Atoi(this.Input().Get("options"))

	if options == 0 {
		// [query_type == 0  GetActivityItemById]
		// [query_type == 1  GetActivityItemByFamilyId]
		// [query_type == 2  GetActivityItemByActivityId]
		query_type, _ := strconv.Atoi(this.Input().Get("query_type"))

		if query_type == 0 {
			activityItemId, _ := strconv.ParseInt(this.Input().Get("activityItemId"), 10, 64)
			activityItem, err := models.GetActivityItemById(activityItemId)
			if err != nil {
				this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			} else {
				this.Data["json"] = map[string]interface{}{"status": 200, "activityItem": activityItem, "time": time.Now().Format("2006-01-02 15:04:05")}
			}
			this.ServeJSON()
			return
		}

		if query_type == 1 {
			familyId, _ := strconv.ParseInt(this.Input().Get("familyId"), 10, 64)
			activityItems, err := models.GetActivityItemByFamilyId(familyId)
			if err != nil {
				this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			} else {
				this.Data["json"] = map[string]interface{}{"status": 200, "activityItems": activityItems, "time": time.Now().Format("2006-01-02 15:04:05")}
			}
			this.ServeJSON()
			return
		}

		if query_type == 2 {
			activityId, _ := strconv.ParseInt(this.Input().Get("activityId"), 10, 64)
			activityItems, err := models.GetActivityItemByActivityId(activityId)
			if err != nil {
				this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			} else {
				this.Data["json"] = map[string]interface{}{"status": 200, "activityItems": activityItems, "time": time.Now().Format("2006-01-02 15:04:05")}
			}
			this.ServeJSON()
			return
		}

	}

	if options == 1 {
		familyId, _ := strconv.ParseInt(this.Input().Get("familyId"), 10, 64)
		activityId, _ := strconv.ParseInt(this.Input().Get("activityId"), 10, 64)
		title := this.Input().Get("title")
		desc := this.Input().Get("desc")
		createTime := this.Input().Get("createTime")
		num, _ := strconv.Atoi(this.Input().Get("num"))

		activityItemId, err := models.AddActivityItem(familyId, activityId, title, desc, createTime, num)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		} else {
			this.Data["json"] = map[string]interface{}{"status": 200, "activityItemId": activityItemId, "time": time.Now().Format("2006-01-02 15:04:05")}
		}
		this.ServeJSON()
		return
	}

	if options == 2 {
		activityItemId, _ := strconv.ParseInt(this.Input().Get("activityItemId"), 10, 64)
		err := models.DelActicityItemById(activityItemId)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		} else {
			this.Data["json"] = map[string]interface{}{"status": 200, "message": "activityItem del success", "time": time.Now().Format("2006-01-02 15:04:05")}
		}
		this.ServeJSON()
		return
	}

	if options == 3 {
		activityItemId, _ := strconv.ParseInt(this.Input().Get("activityItemId"), 10, 64)
		familyId, _ := strconv.ParseInt(this.Input().Get("familyId"), 10, 64)
		activityId, _ := strconv.ParseInt(this.Input().Get("activityId"), 10, 64)
		title := this.Input().Get("activityId")
		desc := this.Input().Get("desc")
		err := models.UpdateActivityItem(activityItemId, familyId, activityId, title, desc)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		} else {
			this.Data["json"] = map[string]interface{}{"status": 200, "message": "activityItem update success", "time": time.Now().Format("2006-01-02 15:04:05")}
		}
		this.ServeJSON()
		return
	}

	this.Data["json"] = map[string]interface{}{"status": 400, "message": "options error", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return
}
