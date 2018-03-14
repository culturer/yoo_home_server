package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//家庭基本信息
type TFamily struct {
	Id                  int64
	FamilyName          string
	FamilyNotifyTitle   string
	FamilyNotifyContent string
	FamilyNotifyTime    string
	CreatedTime         string
	Msg                 string
}

func AddFamily(family *TFamily) (int64, error) {
	o := orm.NewOrm()
	familyId, err := o.Insert(family)
	return familyId, err
}

func GetFamilyById(familyId int64) (*TFamily, error) {
	o := orm.NewOrm()
	family := new(TFamily)
	qs := o.QueryTable("t_family")
	err := qs.Filter("id", familyId).One(family)
	return family, err
}

func DelFamilyById(familyId int64) error {
	o := orm.NewOrm()
	family := &TFamily{Id: familyId}
	_, err := o.Delete(family)
	return err
}

func UpdateFamily(family *TFamily) error {
	o := orm.NewOrm()
	_, err := o.Update(family)
	return err
}
