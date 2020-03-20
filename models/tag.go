package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(name string) (tags []Tag) {
	query := db.Select("id,name,state,created_by,created_on")
	if name != "" {
		query = query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	query.Find(&tags)
	return
}

func GetTag(id int) (tag Tag) {
	db.Where("id = ?", id).First(&tag)

	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func DeleteTag(id int) bool {
	rowsAffected := db.Where("id = ?", id).Delete(Tag{}).RowsAffected
	if rowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func EditTag(id int, data interface{}) bool {
	affected := db.Model(&Tag{}).Where("id = ?", id).Updates(data).RowsAffected
	if affected > 0 {
		return true
	} else {
		return false
	}
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
