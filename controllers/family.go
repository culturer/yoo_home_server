package controllers

import (
	// "github.com/astaxie/beego"
	// "github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
	"yooplus_indication/models"
)

type FamilyController struct {
	BaseController
}

func (this *FamilyController) Get() {
	this.TplName = "family_test.html"
}

//////////////////////////////////////////////////////////////////////
//																	//
//							家庭信息				                //
//																	//
//////////////////////////////////////////////////////////////////////

//options == 0 查询数据
//options == 1 新增数据
//options == 2 删除数据
//options == 3 更新数据
//options == 4 删除用户数据

func (this *FamilyController) Post() {

	// valide, userId, err := this.indicateToken()
	// if err != nil {
	// 	this.Data["json"] = map[string]interface{}{"status": 400, "message": "token is not exist", "time": time.Now().Format("2006-01-02 15:04:05")}
	// 	this.ServeJSON()
	// 	return
	// }
	// if !valide {
	// 	this.Data["json"] = map[string]interface{}{"status": 400, "message": "token is timeout", "time": time.Now().Format("2006-01-02 15:04:05")}
	// 	this.ServeJSON()
	// 	return
	// }

	userId, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
	familyId, _ := strconv.ParseInt(this.Input().Get("familyId"), 10, 64)
	options, _ := strconv.ParseInt(this.Input().Get("options"), 10, 64)
	if options == 0 {
		//查询数据
		family, err := this.getFamily(familyId)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "getFamily error", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		users, err := this.getFamilyUsers(familyId)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "get family users error", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "family": family, "users": users, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	if options == 1 {
		//新增数据
		familyName := this.Input().Get("familyName")
		family := &models.TFamily{FamilyName: familyName, FamilyNotifyTitle: "welcome", FamilyNotifyTime: time.Now().Format("2006-01-02 15:04:05"), CreatedTime: time.Now().Format("2006-01-02 15:04:05")}
		familyId, err := this.addFamily(family, userId)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "addFamily error", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		family, err = this.getFamily(familyId)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "addFamily error", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		this.Data["json"] = map[string]interface{}{"status": 200, "family": family, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	if options == 2 {
		//删除数据
		err := this.delFamily(familyId)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "delFamily error", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "message": "delFamily success", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	if options == 3 {
		//修改数据
		familyName := this.Input().Get("familyName")
		familyNotifyTitle := this.Input().Get("familyNotifyTitle")
		familyNotifyContent := this.Input().Get("familyNotifyContent")
		msg := this.Input().Get("msg")

		family, err := this.getFamily(familyId)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "updateFamily error,family is not exist", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		family.FamilyName = familyName
		family.FamilyNotifyTitle = familyNotifyTitle
		family.FamilyNotifyContent = familyNotifyContent
		family.FamilyNotifyTime = time.Now().Format("2006-01-02 15:04:05")
		family.Msg = msg
		err = this.updateFamily(family)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "updateFamily error,updated exist error", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "message": "updateFamily success", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	if options == 4 {
		err := this.delUserFamily(userId)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "delUserFamily error", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "message": "updateFamily success", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

}

//////////////////////////////////////////////////////////////////////
//																	//
//				将数据库操作再封装一层方便替换model层               //
//																	//
//////////////////////////////////////////////////////////////////////

func (this *FamilyController) getFamily(familyId int64) (*models.TFamily, error) {
	family, err := models.GetFamilyById(familyId)
	return family, err
}

func (this *FamilyController) getFamilyUsers(familyId int64) ([]*models.TUser, error) {
	users, err := models.GetUserByFamilyId(familyId)
	return users, err
}

func (this *FamilyController) addFamily(family *models.TFamily, userId int64) (int64, error) {
	familyId, err := models.AddFamily(family)
	return familyId, err
}

func (this *FamilyController) delFamily(familyId int64) error {
	err := models.DelFamilyById(familyId)
	return err
}

func (this *FamilyController) updateFamily(family *models.TFamily) error {
	err := models.UpdateFamily(family)
	return err
}

func (this *FamilyController) delUserFamily(userId int64) error {
	user, err := models.GetUserById(userId)
	if err != nil {
		return err
	}
	user.FamilyId = -1
	user.FamilyName = ""
	err = models.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}
