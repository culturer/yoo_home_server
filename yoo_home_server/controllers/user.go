package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
	"yoo_home_server/models"
)

type UserController struct {
	BaseController
}

//测试页面
func (this *UserController) Get() {
	this.TplName = "user_test.html"
}

func (this *UserController) Post() {
	userId, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
	options, _ := strconv.ParseInt(this.Input().Get("options"), 10, 64)
	beego.Info(options)
	// [options == 0  修改用户名]
	// [options == 1  修改密码]
	// [options == 2  修改电话号码]
	// [options == 3  修改邮箱]
	// [options == 4  修改头像]
	// [options == 5  修改家庭]
	// [options == 6  修改伴侣]
	// [options == 7  修改签名]
	// [options == 8  修改权限]

	// [options == 0  修改用户名]
	if options == 0 {
		username := this.Input().Get("username")
		err := models.MdfyUsername(username, userId)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		} else {
			this.Data["json"] = map[string]interface{}{"status": 200, "message": "修改用户名成功", "time": time.Now().Format("2006-01-02 15:04:05")}
		}
		this.ServeJSON()
		return
	}

	// [options == 1  修改密码]
	if options == 1 {
		password := this.Input().Get("password")
		err := models.MdfyPassword(password, userId)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		} else {
			this.Data["json"] = map[string]interface{}{"status": 200, "message": "修改密码成功", "time": time.Now().Format("2006-01-02 15:04:05")}
		}
		this.ServeJSON()
		return
	}

	// [options == 2  修改电话号码]
	if options == 2 {
		tel := this.Input().Get("tel")
		err := models.MdfyTel(tel, userId)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		} else {
			this.Data["json"] = map[string]interface{}{"status": 200, "message": "修改电话成功", "time": time.Now().Format("2006-01-02 15:04:05")}
		}
		this.ServeJSON()
		return
	}

	// [options == 3  修改邮箱]
	if options == 3 {
		email := this.Input().Get("email")
		err := models.MdfyEmail(email, userId)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		} else {
			this.Data["json"] = map[string]interface{}{"status": 200, "message": "修改邮箱成功", "time": time.Now().Format("2006-01-02 15:04:05")}
		}
		this.ServeJSON()
		return
	}

	// [options == 4  修改头像]
	if options == 4 {
		icon := this.Input().Get("icon")
		err := models.MdfyIcon(icon, userId)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		} else {
			this.Data["json"] = map[string]interface{}{"status": 200, "message": "修改头像成功", "time": time.Now().Format("2006-01-02 15:04:05")}
		}
		this.ServeJSON()
		return
	}

	// [options == 5  修改家庭]
	if options == 5 {
		familyId, _ := strconv.ParseInt(this.Input().Get("familyId"), 10, 64)
		family, err := models.GetFamilyById(familyId)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		err = models.MdfyFamily(familyId, family.FamilyName, userId)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		} else {
			this.Data["json"] = map[string]interface{}{"status": 200, "message": "修改家庭成功", "time": time.Now().Format("2006-01-02 15:04:05")}
		}
		this.ServeJSON()
		return
	}

	// [options == 6  修改伴侣]
	if options == 6 {
		mateId, _ := strconv.ParseInt(this.Input().Get("mateId"), 10, 64)
		err := models.MdfyMate(mateId, userId)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		} else {
			this.Data["json"] = map[string]interface{}{"status": 200, "message": "修改伴侣成功", "time": time.Now().Format("2006-01-02 15:04:05")}
		}
		this.ServeJSON()
		return
	}

	// [options == 7  修改签名]
	if options == 7 {
		msg := this.Input().Get("msg")
		err := models.MdfyMsg(msg, userId)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		} else {
			this.Data["json"] = map[string]interface{}{"status": 200, "message": "修改签名成功", "time": time.Now().Format("2006-01-02 15:04:05")}
		}
		this.ServeJSON()
		return
	}

	// [options == 8  修改权限]
	if options == 8 {
		per, _ := strconv.Atoi(this.Input().Get("per"))
		err := models.MdfyPer(per, userId)
		if err != nil {
			beego.Info(err.Error())
			this.Data["json"] = map[string]interface{}{"status": 400, "message": err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
		} else {
			this.Data["json"] = map[string]interface{}{"status": 200, "message": "修改权限成功", "time": time.Now().Format("2006-01-02 15:04:05")}
		}
		this.ServeJSON()
		return
	}
}
