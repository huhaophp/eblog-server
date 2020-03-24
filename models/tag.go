package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	ID         int    `gorm:"primary_key" json:"id"`
	CreatedOn  int    `json:"created_on"`
	ModifiedOn int    `json:"modified_on"`
	Name       string `json:"name"`
}

// GetTags 获取所有标签
func GetTags(name string) (tags []Tag) {
	query := db.Select("id,name,created_on,modified_on")
	if name != "" {
		query = query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	query.Order("id DESC").Find(&tags)
	return
}

// GetTag 根据条件标签
func GetTag(where *Tag) (tag Tag) {
	db.Where(where).First(&tag)
	return
}

// AddTag 创建标签
func AddTag(tag *Tag) int64 {
	return db.Create(tag).RowsAffected
}

// EditTag 修改标签
func EditTag(id int, name string) bool {
	db.Model(&Tag{}).Where("id = ?", id).Update("name", name)
	return true
}

// DelTag 标签删除
func DelTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

// 根据ID查询标签
func GetTagsByIds(ids []int) (tags []Tag) {
	db.Select("id,name").Where("id in (?)", ids).Find(&tags)
	return
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
