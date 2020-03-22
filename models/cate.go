package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Cate struct {
	ID         int    `gorm:"primary_key" json:"id"`
	Name       string `json:"name"`
	State      int    `json:"state"`
	CreatedOn  int    `json:"created_on"`
	ModifiedOn int    `json:"modified_on"`
}

// 获取所有栏目
func GetCates(name string) (cates []Cate) {
	query := db.Select("id,name,state,created_on,modified_on")
	if name != "" {
		query = query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	query.Find(&cates)
	return
}

// GetCate 获取单个栏目
func GetCate(where *Cate) (cate Cate) {
	db.Where(where).First(&cate)
	return
}

// AddCate 栏目创建
func AddCate(cate *Cate) int64 {
	return db.Create(cate).RowsAffected
}

// EditCate 栏目编辑
func EditCate(id int, name string, state int) bool {
	db.Model(&Cate{}).Where("id = ?", id).Updates(Cate{Name: name, State: state})
	return true
}

// DelCate 栏目删除
func DelCate(id int) bool {
	db.Where("id = ?", id).Delete(&Cate{})
	return true
}

func (tag *Cate) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (tag *Cate) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
