package controllers

import (
	"github.com/astaxie/beego"
	"io"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"
	"yoo_home_server/models"
)

type FileController struct {
	beego.Controller
}

//文件下载
func (this *FileController) Get() {
	this.TplName = "file_test.html"

}

//文件上传
func (this *FileController) Post() {

	options, _ := strconv.ParseInt(this.Input().Get("options"), 10, 64)
	albumId, _ := strconv.ParseInt(this.Input().Get("albumId"), 10, 64)
	familyId, _ := strconv.ParseInt(this.Input().Get("familyId"), 10, 64)
	userId := this.Input().Get("userId")

	beego.Info(options)
	beego.Info(userId)

	//创建用户目录
	err := os.MkdirAll("photos/"+userId, os.ModePerm)
	if err != nil {
		beego.Error(err)
	}

	if options == 0 {
		// 获取附件
		_, fh, ee := this.GetFile("attachment")
		if ee != nil {
			beego.Error(ee)
			this.Data["json"] = map[string]interface{}{"status": 400, "message": ee.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		var attachment string
		if fh != nil {
			//保存附件
			attachment = fh.Filename
			beego.Info(attachment)
			myPath := path.Join("photos/"+userId, attachment)
			beego.Info(myPath)
			err := this.SaveToFile("attachment", myPath)

			if err != nil {
				beego.Error(err)
				this.Data["json"] = map[string]interface{}{"status": 400, "message": "upload fail", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			//数据信息存数据库
			mUserId, err := strconv.ParseInt(userId, 10, 64)
			if err != nil {
				beego.Info(err)
			}

			photo, err := this.addPhoto(albumId, mUserId, familyId, attachment, myPath, time.Now().Format("2006-01-02 15:04:05"))

			if err != nil {
				beego.Error(err)
				this.Data["json"] = map[string]interface{}{"status": 400, "message": "add photo fail", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}

			this.Data["json"] = map[string]interface{}{"status": 200, "photo": photo, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return

		}

	}

	// 下载接口
	if options == 1 {
		filePath, err := url.QueryUnescape(this.Ctx.Request.RequestURI[1:])
		if err != nil {
			this.Ctx.WriteString(err.Error())
			return
		}
		f, err := os.Open(filePath)
		if err != nil {
			this.Ctx.WriteString(err.Error())
			return
		}
		defer f.Close()
		_, err = io.Copy(this.Ctx.ResponseWriter, f)
		if err != nil {
			this.Ctx.WriteString(err.Error())
			return
		}
	}

	this.Data["json"] = map[string]interface{}{"status": 400, "message": "options error", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

func (this *FileController) addPhoto(albumId, userId, familyId int64, fileName, fileUrl, createTime string) (*models.TPhoto, error) {
	photo := &models.TPhoto{AlbumId: albumId, UserId: userId, FamilyId: familyId, FileName: fileName, FileUrl: fileUrl, CreateTime: createTime}
	photoId, err := models.AddPhoto(photo)
	photo.Id = photoId
	return photo, err
}
