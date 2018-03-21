package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Get() {

	this.TplName = "base_test.tpl"
}

func (this *BaseController) Post() {
	isValid, _, err := this.indicateToken()
	if err != nil || !isValid {
		this.Data["json"] = map[string]interface{}{"status": 400, "message": "自动登陆失败", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{"status": 200, "message": "自动登陆成功", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return
}

// ParseToken parse JWT token in http header.
func (c *BaseController) ParseToken() (t *jwt.Token, e error) {

	authString := c.Ctx.Input.Header("Authorization")
	beego.Debug("AuthString:", authString)

	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		beego.Error("AuthString invalid:", authString)
		return nil, errors.New("errInputData")
	}
	tokenString := kv[1]

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("yooplus"), nil
	})
	if err != nil {
		beego.Error("Parse token:", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// That‘s not even a token
				return nil, errors.New("errInputData")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return nil, errors.New("errExpired")
			} else {
				// Couldn‘t handle this token
				return nil, errors.New("errInputData")
			}
		} else {
			// Couldn‘t handle this token
			return nil, errors.New("errInputData")
		}
	}
	if !token.Valid {
		beego.Error("Token invalid:", tokenString)
		return nil, errors.New("errInputData")
	}
	beego.Debug("Token:", token)

	return token, nil
}

//验证token是否有效
func (this *BaseController) indicateToken() (bool, int64, error) {

	token, err := this.ParseToken()
	if err != nil {
		return false, -1, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, -1, nil
	}
	var userId int64 = claims["userId"].(int64)
	return true, userId, nil
}
