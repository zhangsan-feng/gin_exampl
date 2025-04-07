package resources_entity

import "gin_exampl/global"

type ResourcesEntity struct {
}

func (ResourcesEntity) ResourcesQuery() {
	global.DBConn.Table("resources")
}

func (ResourcesEntity) ResourcesAdd() {
	global.DBConn.Table("resources")
}

func (ResourcesEntity) ResourcesEdit() {
	global.DBConn.Table("resources")
}

func (ResourcesEntity) ResourcesDelete() {
	global.DBConn.Table("resources")
}
