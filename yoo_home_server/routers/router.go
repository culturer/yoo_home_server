package routers

import (
	"github.com/astaxie/beego"
	"yoo_home_server/controllers"
)

func init() {
	beego.Router("/", &controllers.BaseController{})
	//用户
	beego.Router("/user", &controllers.UserController{})
	//登录
	beego.Router("/login", &controllers.LoginController{})
	//注册
	beego.Router("/register", &controllers.RegisterController{})
	//相册
	beego.Router("/albums", &controllers.AlbumController{})
	//图片上传
	beego.Router("/photos", &controllers.PhotoController{})
	//家庭
	beego.Router("/family", &controllers.FamilyController{})
	//活动
	beego.Router("/activities", &controllers.ActivityController{})
	//活动项
	beego.Router("/activityitems", &controllers.ActivityItemController{})
	//日程安排
	beego.Router("/arrangement", &controllers.ArrangementController{})
	//文章
	beego.Router("/article", &controllers.ArticleController{})
	//文件服务
	beego.Router("/files", &controllers.FileController{})
	//验证码
	beego.Router("/captcha", &controllers.CaptchaController{})

}
