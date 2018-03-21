package controllers

import (
	// "github.com/astaxie/beego"
	// "github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
	"yoo_home_server/models"
)

//日程安排接口
type ArrangementController struct {
	BaseController
}

//测试页面
func (this *ArrangementController) Get() {
	this.TplName = "arrangement_test.html"
}

//////////////////////////////////////////////////////////////////////
//																	//
//						获取Arrangement接口			                //
//																	//
//////////////////////////////////////////////////////////////////////

// [options == 0  查询]
// [options == 1  增加]
// [options == 2  删除]
// [options == 3  修改]

func (this *ArrangementController) Post() {
	userId, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
	options, _ := strconv.ParseInt(this.Input().Get("options"), 10, 64)

	//查询
	if options == 0 {
		arrangements, err := this.getArrangements(userId)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "getArrangements fail", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "arrangements": arrangements, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	//增加
	if options == 1 {
		desc := this.Input().Get("desc")
		createTime := this.Input().Get("createTime")
		arrangement := &models.TUserArrangement{UserId: userId, Desc: desc, CreateTime: createTime}
		arrangementId, err := this.addArrangement(arrangement)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "addArrangements fail", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "arrangementId": arrangementId, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	if options == 2 {
		arrangementId, _ := strconv.ParseInt(this.Input().Get("arrangementId"), 10, 64)
		err := this.delArrangement(arrangementId)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "delArrangements fail", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "message": "delarrangement success", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	if options == 3 {
		arrangementId, _ := strconv.ParseInt(this.Input().Get("arrangementId"), 10, 64)
		desc := this.Input().Get("desc")
		err := this.updateArrangement(arrangementId, desc)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "updateArrangements fail", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "message": "updatearrangement success", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	this.Data["json"] = map[string]interface{}{"status": 400, "message": " option fail", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return
}

//////////////////////////////////////////////////////////////////////
//																	//
//				将数据库操作再封装一层方便替换model层               //
//																	//
//////////////////////////////////////////////////////////////////////

func (this *ArrangementController) getArrangements(userId int64) ([]*models.TUserArrangement, error) {
	arrangements, err := models.GetUserArrangementByUserId(userId)
	return arrangements, err
}
func (this *ArrangementController) addArrangement(arrangement *models.TUserArrangement) (int64, error) {
	arrangementId, err := models.AddUserArrangement(arrangement)
	return arrangementId, err
}
func (this *ArrangementController) delArrangement(arrangementId int64) error {
	arrangement := &models.TUserArrangement{Id: arrangementId}
	err := models.DelUserArrangement(arrangement)
	return err
}
func (this *ArrangementController) updateArrangement(arrangementId int64, desc string) error {
	arrangement, err := models.GetUserArrangementById(arrangementId)
	if err != nil {
		return err
	}
	arrangement.Desc = desc
	err = models.UpdateUserArrangement(arrangement)
	return err
}
