package controllers

import (
	// "github.com/astaxie/beego"
	// "github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
	"yoo_home_server/models"
)

type AlbumController struct {
	BaseController
}

//测试页面
func (this *AlbumController) Get() {
	this.TplName = "albumItem_test.html"
}

//////////////////////////////////////////////////////////////////////
//																	//
//							相册API接口				                //
//																	//
//////////////////////////////////////////////////////////////////////

// [options == 0  查询]
// [options == 1  增加]
// [options == 2  删除]
// [options == 3  修改]

func (this *AlbumController) Post() {

	var userId int64
	var familyId int64
	var albumType bool
	var albums []*models.TAlbumItem
	var err error

	options, _ := strconv.ParseInt(this.Input().Get("options"), 10, 64)

	if options == 0 {
		albumType, _ = strconv.ParseBool(this.Input().Get("albumType"))

		//albumType = false -- FamilyAlbum
		//albumType = true  -- UserAlbum
		if albumType {
			userId, _ = strconv.ParseInt(this.Input().Get("userId"), 10, 64)
			albums, err = this.getUserAlbums(userId)
			if err != nil {
				this.Data["json"] = map[string]interface{}{"status": 400, "message": "getUserAlbums fail", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
		} else {
			familyId, _ = strconv.ParseInt(this.Input().Get("familyId"), 10, 64)
			albums, err = this.getFamilyAlbum(familyId)
			if err != nil {
				this.Data["json"] = map[string]interface{}{"status": 400, "message": "getFamilyAlbum fail", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
		}

		this.Data["json"] = map[string]interface{}{"status": 200, "albums": albums, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}
	if options == 1 {
		albumType, _ := strconv.ParseBool(this.Input().Get("albumType"))
		name := this.Input().Get("albumItemName")
		if albumType {
			id, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
			albumItemId, err := this.addAlbumItem(albumType, name, id)
			if err != nil {
				this.Data["json"] = map[string]interface{}{"status": 400, "message": "addAlbumItem fail", "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "albumItemId": albumItemId, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		} else {
			id, _ := strconv.ParseInt(this.Input().Get("familyId"), 10, 64)
			albumItemId, err := this.addAlbumItem(albumType, name, id)
			if err != nil {
				this.Data["json"] = map[string]interface{}{"status": 400, "albumItemId": albumItemId, "time": time.Now().Format("2006-01-02 15:04:05")}
				this.ServeJSON()
				return
			}
			this.Data["json"] = map[string]interface{}{"status": 200, "albumItemId": albumItemId, "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

	}

	if options == 2 {
		albumItemId, _ := strconv.ParseInt(this.Input().Get("albumItemId"), 10, 64)
		err := this.delAlbumItem(albumItemId)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "delAlbumItem fail", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "message": "delAlbumItem success", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	if options == 3 {
		albumItemId, _ := strconv.ParseInt(this.Input().Get("albumItemId"), 10, 64)
		albumItemName := this.Input().Get("albumItemName")
		err := this.updateAlbumItem(albumItemId, albumItemName)
		if err != nil {
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "updateAlbumItem fail", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}

		this.Data["json"] = map[string]interface{}{"status": 200, "message": "updateAlbumItem success", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	this.Data["json"] = map[string]interface{}{"status": 400, "message": "options fail", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return
}

//////////////////////////////////////////////////////////////////////
//																	//
//				将数据库操作再封装一层方便替换model层               //
//																	//
//////////////////////////////////////////////////////////////////////

func (this *AlbumController) getUserAlbums(userId int64) ([]*models.TAlbumItem, error) {
	albums, err := models.GetUserAlbumItem(userId)
	return albums, err
}
func (this *AlbumController) getFamilyAlbum(familyId int64) ([]*models.TAlbumItem, error) {
	albums, err := models.GetFamilyAlbumItem(familyId)
	return albums, err
}
func (this *AlbumController) addAlbumItem(alubmType bool, name string, id int64) (int64, error) {
	if alubmType {
		albumItem := &models.TAlbumItem{UserId: id, Name: name, CreateTime: time.Now().Format("2006-01-02 15:04:05")}
		albumItemId, err := models.AddAlbumItem(albumItem)
		return albumItemId, err
	} else {
		albumItem := &models.TAlbumItem{FamilyId: id, Name: name, CreateTime: time.Now().Format("2006-01-02 15:04:05")}
		albumItemId, err := models.AddAlbumItem(albumItem)
		return albumItemId, err
	}
}
func (this *AlbumController) delAlbumItem(albumItemId int64) error {
	albumItem := &models.TAlbumItem{Id: albumItemId}
	err := models.DelAlbumItem(albumItem)
	return err
}
func (this *AlbumController) updateAlbumItem(albumItemId int64, albumItemName string) error {
	albumItem, err := models.GetAlbumItemById(albumItemId)
	if err != nil {
		return err
	}
	albumItem.Name = albumItemName
	err = models.UpdateAlbumItem(albumItem)
	return err
}
