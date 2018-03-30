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

func MdfyUsername(username string, userId int64) error {
	user, err := GetUserById(userId)
	if err != nil {
		return nil
	}
	user.Username = username
	o := orm.NewOrm()
	_, err = o.Update(user)
	return err
}

func MdfyPassword(password string, userId int64) error {
	user, err := GetUserById(userId)
	if err != nil {
		return nil
	}
	user.Password = password
	o := orm.NewOrm()
	_, err = o.Update(user)
	return err
}

func MdfyTel(tel string, userId int64) error {
	user, err := GetUserById(userId)
	if err != nil {
		return nil
	}
	user.Tel = tel
	o := orm.NewOrm()
	_, err = o.Update(user)
	return err
}

func MdfyEmail(email string, userId int64) error {
	user, err := GetUserById(userId)
	if err != nil {
		return nil
	}
	user.Email = email
	o := orm.NewOrm()
	_, err = o.Update(user)
	return err
}

func MdfyIcon(icon string, userId int64) error {
	user, err := GetUserById(userId)
	if err != nil {
		return nil
	}
	user.Icon = icon
	o := orm.NewOrm()
	_, err = o.Update(user)
	return err
}

func MdfyFamily(familyId int64, familyName string, userId int64) error {
	user, err := GetUserById(userId)
	if err != nil {
		return nil
	}
	user.FamilyId = familyId
	user.FamilyName = familyName
	o := orm.NewOrm()
	_, err = o.Update(user)
	return err
}

func MdfyMate(mateId int64, userId int64) error {
	user, err := GetUserById(userId)
	if err != nil {
		return nil
	}
	user.MateId = mateId
	o := orm.NewOrm()
	_, err = o.Update(user)
	return err
}

func MdfyMsg(msg string, userId int64) error {
	user, err := GetUserById(userId)
	if err != nil {
		return nil
	}
	user.Msg = msg
	o := orm.NewOrm()
	_, err = o.Update(user)
	return err
}

func MdfyPer(per int, userId int64) error {
	user, err := GetUserById(userId)
	if err != nil {
		return nil
	}
	user.Permission = per
	o := orm.NewOrm()
	_, err = o.Update(user)
	return err
}
