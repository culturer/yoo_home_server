package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func RegiesterDB() {

	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:78901214@tcp(127.0.0.1:3306)/yooplus_test1?charset=utf8")
	//注册model
	orm.RegisterModel(new(TUser), new(TUserArrangement))
	orm.RegisterModel(new(TFamily), new(TFamilyActivity))
	orm.RegisterModel(new(TAlbumItem), new(TPhoto), new(TActivityItem))
	orm.RegisterModel(new(TAddress))
	orm.RegisterModel(new(TArticle), new(TComment))

}
