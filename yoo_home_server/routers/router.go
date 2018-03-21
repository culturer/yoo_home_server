package routers

import (
	"github.com/astaxie/beego"
	"yoo_home_server/controllers"
)

func init() {
	beego.Router("/", &controllers.BaseController{})
	//业务服务
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/albums", &controllers.AlbumController{})
	beego.Router("/photos", &controllers.PhotoController{})
	beego.Router("/family", &controllers.FamilyController{})
	beego.Router("/activities", &controllers.ActivityController{})
	beego.Router("/activityitems", &controllers.ActivityItemController{})
	beego.Router("/arrangement", &controllers.ArrangementController{})
	beego.Router("/article", &controllers.ArticleController{})
	//文件服务
	beego.Router("/files", &controllers.FileController{})
	//验证码
	beego.Router("/captcha", &controllers.CaptchaController{})

}
