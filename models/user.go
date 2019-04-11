package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
)

type User struct {
	Id       int
	Username string `orm:"size(32)" `
	Password string `orm:"size(32)" `
	Email    string `orm:"size(32)"`
	Role     int    `orm:"default(1)"`

	//Phone    string `orm:"size(11)"`\
	//Level    int    `orm:"size(2)"`
	//Vip      bool
	//Regestertime time.Time
	//Birthday time.Time
}

func UserIsRegistered(email string) bool {

	return db.QueryTable("user").Filter("email", email).Exist()
}

func CheckUser(email, password string) (User, error) {

	user := User{}
	if len(strings.TrimSpace(email)) > 0 {
		user.Email = email
	}
	if len(strings.TrimSpace(password)) > 0 {
		user.Password = password
	}
	if err := db.Read(&user, "email", "password"); err != nil {
		logs.Error("checkuser error", err.Error())
		return user, err
	}
	logs.Info("check user successfully.", fmt.Sprintf("%#v", user))
	return user, nil
}

func AddUser(user *User) (bool, error) {
	id, err := db.Insert(user)
	if err != nil {
		logs.Error("add user failed-->", err.Error())
		return false, err
	}
	logs.Debug("add user successfully,inserted user id :" + strconv.Itoa(int(id)))
	return true, nil
}

func UpdateUser(user *User) (bool, error) {
	id, err := db.Update(user)
	if err != nil {
		logs.Error("update user failed-->", err.Error())
		return false, err
	}
	logs.Debug("update user successfully,inserted user id :" + strconv.Itoa(int(id)))
	return true, nil
}

func DelUser(id int) (bool, error) {
	index, err := db.Delete(GetUser(id))
	if err != nil {
		logs.Error("delete user failed-->", err.Error())
		return false, err
	}
	logs.Debug("delete user successfully,inserted user id :" + strconv.Itoa(int(index)))
	return true, nil
}

func GetUsers() []User {
	var maps []orm.Params
	_, err := db.Raw("select * from user").Values(&maps)
	if err != nil {
		logs.Error("this is occured error when get all users-->", err.Error())
		return nil
	}
	var users = make([]User, len(maps))
	for index, item := range maps {
		id, _ := strconv.ParseInt(item["id"].(string), 10, 32)
		users[index] = User{Id: int(id),
			Username: item["username"].(string),
			Password: item["password"].(string),
			Email:    item["email"].(string)}
	}
	return users

}

func GetUser(id int) *User {
	user := User{Id: id}
	if err := db.Read(&user); err != nil {
		logs.Error("there is occured the error when get user-->", err.Error())
		return nil
	}
	logs.Debug("get user successfully ", fmt.Sprintf("%#v", user))
	return &user
}
