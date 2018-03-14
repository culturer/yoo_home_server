package models

import (
// "github.com/astaxie/beego/orm"
// _ "github.com/go-sql-driver/mysql"
)

//个人信息的关系表
type TAddress struct {
	Id       int64
	Country  string
	Province string
	City     string
	Reginon  string
	Street   string
	desc     string
}
