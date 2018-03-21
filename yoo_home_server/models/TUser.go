package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//用户基本表
type TUser struct {
	//用户基本信息
	Id int64
	// 用户名
	Username string
	// 密码
	Password string
	// 真实姓名
	RealName string
	// 性别
	Sex bool
	// 身份证号
	Uid string
	//生日
	Birth string
	//电话
	Tel string `orm:"index"`
	//邮箱
	Email string
	//头像路径
	Icon string
	//签名
	NMsg string
	//关系表
	RelationId int64
	//家庭表
	FamilyId int64
	//家庭名称
	FamilyName string
	//父亲Id
	FatherId int64
	//母亲Id
	MotherId int64
	//配偶Id
	MateId int64
	//创建时间
	CreatedTime string
	//最后登录时间
	LoginTime string
	//预留字段-备注
	Msg string
	//预留字段-权限
	Permission int
}

func AddUser(user *TUser) (int64, error) {
	o := orm.NewOrm()
	userId, err := o.Insert(user)
	return userId, err
}

func UpdateUser(user *TUser) error {
	o := orm.NewOrm()
	_, err := o.Update(user)
	return err
}

func GetUserByTel(tel string) (*TUser, error) {
	o := orm.NewOrm()
	user := new(TUser)
	qs := o.QueryTable("t_user")
	err := qs.Filter("tel", tel).One(user)
	return user, err
}

func GetUserById(userId int64) (*TUser, error) {
	o := orm.NewOrm()
	user := new(TUser)
	qs := o.QueryTable("t_user")
	err := qs.Filter("id", userId).One(user)
	return user, err
}

func GetUserByFamilyId(familyId int64) ([]*TUser, error) {
	users := make([]*TUser, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("t_user")
	_, err := qs.Filter("family_id", familyId).All(&users)
	return users, err
}
