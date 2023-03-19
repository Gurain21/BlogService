package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	//嵌入gorm定义的Model
	Model
	
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	DeletedOn  int    `json:"deleted_on"`
	State      int    `json:"state"`
}

func GetTags(pageNum, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return tags
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return count
}

func InsertTag(tag *Tag) bool {
	return db.Create(tag).Error == nil
}

func AddTag(name, createdBy string, state int) bool {
	tag := &Tag{
		Name:      name,
		CreatedBy: createdBy,
		State:     state,
	}
	return db.Create(tag).Error == nil
	
}
func ExistTagByName(name string) bool {
	count := 0
	db.Model(&Tag{}).Where("name = ?", name).Count(&count)
	return count != 0
}

func ExistTagById(id int) bool {
	t := &Tag{}
	db.Model(&Tag{}).First(t, id)
	return t.Id > 0
}

func UpdateTag(maps interface{}) bool {
	
	return db.Model(&Tag{}).Update(maps).Error == nil
}

func DeleteTagById(id int) bool {
	return db.Where("id = ?", id).Delete(&Tag{}).Error == nil
}

/*
对象生命周期
Hook 是在创建、查询、更新、删除等操作之前、之后调用的函数。

如果您已经为模型定义了指定的方法，它会在创建、更新、查询、删除时自动被调用。如果任何回调返回错误，GORM 将停止后续的操作并回滚事务。

钩子方法的函数签名应该是 func(*gorm.DB) error

Hook

钩子方法：
		创建：BeforeSave、BeforeCreate、AfterCreate、AfterSave
		更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave
		删除：BeforeDelete、AfterDelete
		查询：AfterFind

*/

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	
	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	
	return nil
}
