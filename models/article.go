package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	ID int `gorm:"primary_key" json:"id"`

	CateId int  `json:"cate_id"`
	Cate   Cate `json:"cate"`

	Tags []Tag `json:"tags" gorm:"many2many:article_tag;"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Cover      string `json:"cover"`
	Content    string `json:"content"`
	State      int    `json:"state"`
	CreatedOn  int    `json:"created_on"`
	ModifiedOn int    `json:"modified_on"`
}

// 获取文章列表
func GetArticles(where *Article, limit int, offset int) (articles []Article) {
	query := db.Model(Article{}).Preload("Cate").Preload("Tags", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,name,tag_id,article_id")
	})
	if where.Title != "" {
		query = query.Where("title Like ?", fmt.Sprintf("%%%s%%", where.Title))
	}
	if where.State > -1 {
		query = query.Where("state = ?", where.State)
	}
	query.Order("id DESC").Limit(limit).Offset(offset).Find(&articles)
	return
}

// 获取文章数量
func GetArticlesTotal(where *Article) (total int) {
	query := db.Model(Article{})
	if where.Title != "" {
		query = query.Where("title Like ?", fmt.Sprintf("%%%s%%", where.Title))
	}
	if where.State > -1 {
		query = query.Where("state = ?", where.State)
	}
	query.Count(&total)
	return
}

// AddCate 文章创建
func AddArticle(article *Article) error {
	return db.Create(article).Association("Tags").Replace(article.Tags).Error
}

// AddCate 栏目创建
func EditArticle(id int, article *Article) error {
	return db.Model(&Article{}).Where("id = ?", id).Updates(article).Association("Tags").Replace(article.Tags).Error
}

// DelCate 栏目删除
func DelArticle(id int) error {
	return db.Where("id = ?", id).Delete(&Article{}).Error
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}