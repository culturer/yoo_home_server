package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
	"yoo_home_server/models"
	_ "yoo_home_server/routers"
)

func init() {
	models.RegiesterDB()
}

//////////////////////////////////////////////////////////////////////
//																	//
//						APP基本信息后台接口			                //
//																	//
//////////////////////////////////////////////////////////////////////

func main() {

	// //开启orm调试
	orm.Debug = false
	// //自动建表
	orm.RunSyncdb("default", false, true)

	//创建附件目录
	os.Mkdir("photos", os.ModePerm)
	beego.SetStaticPath("photos", "photos")

	beego.Run()

}
