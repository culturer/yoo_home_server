package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"yoo_home_server/models"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.Data["username"] = "123456"
	c.Data["password"] = "123456"
	c.TplName = "register_test.html"
}

//////////////////////////////////////////////////////////////////////
//																	//
//							注册API接口				                //
//																	//
//////////////////////////////////////////////////////////////////////

func (this *RegisterController) Post() {

	//初始化参数
	var err error = nil
	var userId int64 = -1
	var userAlbumId int64 = -1

	//获取数据信息
	password := this.Input().Get("password")
	tel := this.Input().Get("tel")

	//判断该手机号是否已经注册
	user, err := this.getUser(tel)
	if user.Id != 0 {
		beego.Info(user.Id)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": " 注册失败,该手机号已被注册 ", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	if err != nil {
		beego.Info(err.Error())
	}

	familyId, err := this.addFamily("优 家", " 欢迎来到优家 ")
	beego.Info(familyId)
	if err != nil {
		beego.Info(err.Error())
	}

	err = this.addFamilyActivity(familyId, "快来发起第一个家庭活动吧~")
	if err != nil {
		beego.Info(err.Error())
	}

	familyAlbumId, err := this.addFamilyAlbum(familyId)
	beego.Info(familyAlbumId)
	if err != nil {
		beego.Info(err.Error())
	}

	//注册用户
	user = &models.TUser{Password: password, Tel: tel, FamilyId: familyId, CreatedTime: time.Now().Format("2006-01-02 15:04:05")}
	userId, err = this.addUser(user)
	if err != nil {
		beego.Error(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "msg": "register fail -- add user fail ", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	if userId != -1 {
		//新建UserAlbum
		userAlbumId, err = this.addUserAlbum(userId)
		beego.Info(userAlbumId)
		if err != nil {
			beego.Info(err.Error())
		}
		//新建一条userArrangement
		err = this.addUserArrangement(userId, "")
		if err != nil {
			beego.Info(err.Error())
		}

	}

	this.Data["json"] = map[string]interface{}{"status": 200, "msg": "register success ", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

//新建User
func (this *RegisterController) addUser(user *models.TUser) (int64, error) {
	userId, err := models.AddUser(user)
	return userId, err
}

//获取User
func (this *RegisterController) getUser(tel string) (*models.TUser, error) {
	user, err := models.GetUserByTel(tel)
	return user, err
}

//新建UserAlbum
func (this *RegisterController) addUserAlbum(userId int64) (int64, error) {
	albumItem := &models.TAlbumItem{UserId: userId, Name: "我的相册", CreateTime: time.Now().Format("2006-01-02 15:04:05")}
	userAlbumId, err := models.AddAlbumItem(albumItem)
	return userAlbumId, err
}

//新建日程安排
func (this *RegisterController) addUserArrangement(userId int64, desc string) error {
	userArrangement := &models.TUserArrangement{UserId: userId, Desc: "Hello Yoo+ ! Hellow " + desc + " !", CreateTime: time.Now().Format("2006-01-02 15:04:05")}
	_, err := models.AddUserArrangement(userArrangement)
	return err
}

func (this *RegisterController) addFamily(familyName, notfy string) (int64, error) {
	family := &models.TFamily{FamilyNotifyTitle: notfy, FamilyName: familyName, FamilyNotifyTime: time.Now().Format("2006-01-02 15:04:05"), CreatedTime: time.Now().Format("2006-01-02 15:04:05")}
	familyId, err := models.AddFamily(family)
	return familyId, err
}

func (this *RegisterController) addFamilyActivity(familyId int64, desc string) error {
	familyActivity := &models.TFamilyActivity{FamilyId: familyId, Desc: desc, CreateTime: time.Now().Format("2006-01-02 15:04:05")}
	err := models.AddFamilyActivity(familyActivity)
	return err
}

func (this *RegisterController) addFamilyAlbum(familyId int64) (int64, error) {
	albumItem := &models.TAlbumItem{FamilyId: familyId, Name: "我的家庭相册", CreateTime: time.Now().Format("2006-01-02 15:04:05")}
	userAlbumId, err := models.AddAlbumItem(albumItem)
	return userAlbumId, err
}
