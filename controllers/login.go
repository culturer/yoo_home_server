package controllers

import (
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"time"
	"yooplus_indication/models"
)

type LoginController struct {
	BaseController
}

type Relation struct {
}

//测试页面入口
func (c *LoginController) Get() {
	c.Data["username"] = "admin"
	c.Data["password"] = "123456"
	c.TplName = "login_test.html"
}

//////////////////////////////////////////////////////////////////////
//																	//
//							登录API接口				                //
//																	//
//////////////////////////////////////////////////////////////////////

//目前只做了通过账号密码登陆，还需要做通过手机，邮箱登陆

func (this *LoginController) Post() {

	//初始化变量
	var token string = ""
	var userId int64 = -1
	var err error = nil
	var user *models.TUser = nil
	var family *models.TFamily = nil

	var isIndicate bool = false
	isIndicate, userId, err = this.indicateToken()

	if err != nil {
		beego.Error(err)
	}
	if isIndicate {
		//自动登录成功,则不验证账号密码
		beego.Info("auto login success!")
	} else {
		beego.Info("auto login fail!")
		//自动登录失败,验证账号密码是否正确
		tel := this.Input().Get("tel")
		password := this.Input().Get("password")
		var isLogin bool = false
		isLogin, token, userId, err = this.indicateUser(tel, password)
		if err != nil || !isLogin {
			//验证用户失败
			beego.Error(err)
			this.Data["json"] = map[string]interface{}{"status": 400, "message": " 登陆失败,账号不存在或账号,密码错误。 ", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

	}

	//获取用户信息
	beego.Info("用户id:", userId)
	user, err = this.getUser(userId)
	if err != nil {
		beego.Error(err)
		this.Data["json"] = map[string]interface{}{"status": 400, "message": "login failed -- getUser error ", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	//获取家庭信息
	family, err = this.getFamily(user.FamilyId)
	if err != nil {
		beego.Error(err)
	}
	familyUsers, err := this.getFamilyUser(user.FamilyId)
	if err != nil {
		beego.Error(err)
	}

	this.Data["json"] = map[string]interface{}{"status": 200, "message": "login success ", "token": token, "user": user, "family": family, "familyUsers": familyUsers, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return
}

//////////////////////////////////////////////////////////////////////
//																	//
//				将数据库操作再封装一层方便替换model层               //
//																	//
//////////////////////////////////////////////////////////////////////

// 验证账号密码
func (this *LoginController) indicateUser(tel, password string) (bool, string, int64, error) {
	user, err := models.GetUserByTel(tel)
	if err != nil {
		beego.Error(err)
		return false, "", -1, err
	}
	if user.Password == password {
		token, err := this.createToken(user.Id)
		if err != nil {
			return false, "", -1, err
		}
		return true, token, user.Id, nil
	}
	return false, "", -1, nil
}

//获取token
func (this *LoginController) createToken(userId int64) (string, error) {
	// 带权限创建令牌
	claims := make(jwt.MapClaims)
	claims["userId"] = userId
	//20天有效期，过期需要重新登录获取token
	claims["exp"] = time.Now().Add(time.Hour * 480).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用自定义字符串加密 and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte("yooplus"))
	if err != nil {
		beego.Error("jwt.SignedString:", err)
		return "", err
	}
	return tokenString, nil
}

//获取用户信息
func (this *LoginController) getUser(userId int64) (*models.TUser, error) {
	user, err := models.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//获取家庭信息
func (this *LoginController) getFamily(familyId int64) (*models.TFamily, error) {
	family, err := models.GetFamilyById(familyId)
	if err != nil {
		return nil, err
	}
	return family, nil
}

//获取家庭成员
func (this *LoginController) getFamilyUser(familyId int64) ([]*models.TUser, error) {
	users, err := models.GetUserByFamilyId(familyId)
	return users, err
}
