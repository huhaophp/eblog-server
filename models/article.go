package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	ID    int `gorm:"primary_key" json:"id"`
	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`
	Title      string `json:"title"`
	Cover      string `json:"cover"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
	CreatedOn  int    `json:"created_on"`
	ModifiedOn int    `json:"modified_on"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
