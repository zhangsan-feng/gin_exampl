package user_entity

import "gin_exampl/global"

type UserEntity struct {
}

func (UserEntity) UserAdd() {
	global.DBConn.Table("user").Create(UserEntity{})
}

func (UserEntity) UserQuery(page int) {
	global.DBConn.Table("user").Where("id > ", page*10).Scan(UserEntity{})
}

func (UserEntity) UserDelete(userId int) {
	global.DBConn.Table("user").Delete("id", userId)
}

func (UserEntity) UserEdit() {

	global.DBConn.Table("user").Updates(&UserEntity{})
}
