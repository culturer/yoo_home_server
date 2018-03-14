package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
	"yooplus_indication/models"
)

type PhotoController struct {
	BaseController
}

//测试页面
func (this *PhotoController) Get() {
	this.TplName = "test.html"
}

//////////////////////////////////////////////////////////////////////
//																	//
//							照片API接口				                //
//																	//
//////////////////////////////////////////////////////////////////////

func (this *PhotoController) Post() {
	userId, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
	beego.Info(userId)
	photos, err := this.getPhotos(userId)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"status": 400, "message": "getPhotos error", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{"status": 200, "photos": photos, "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

//////////////////////////////////////////////////////////////////////
//																	//
//				将数据库操作再封装一层方便替换model层               //
//																	//
//////////////////////////////////////////////////////////////////////

func (this *PhotoController) getPhotos(userId int64) ([]*models.TPhoto, error) {
	photos, err := models.GetPhotos(userId)
	return photos, err
}
