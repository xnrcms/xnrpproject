package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id 					int
	Email 				string
	Name 				string
	Password 			string
	Type 				int
	Status 				int
	LastIp 				string
	Address 			string
	Device 				string
	AddTime 			int64
	IsAdmin 			int
	Salt 				string
	LastLogin 			int64
}

func (a *User) TableName() string {
	return TableName("user")
}

func UserAdd(a *User) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func UserGetById(id int) (*User, error) {
	user := new(User)
	err := orm.NewOrm().QueryTable(user.TableName()).Filter("id", id).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UserGetByField(field string,value string) (*User, error) {
	user := new(User)
	err := orm.NewOrm().QueryTable(user.TableName()).Filter(field, value).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}