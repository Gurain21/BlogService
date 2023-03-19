package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model
	
	TagId int `json:"tag_id" gorm:"index"`
	Tag   `json:"tag"`
	
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (a *Article) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("CreatedOn", time.Now().Unix())
}
func (a *Article) BeforeUpdate(scope *gorm.Scope) error {
	return scope.SetColumn("ModifiedOn", time.Now().Unix())
}

func ExistArticleById(id int) bool {
	var article Article
	db.Model(&Article{}).Where("id = ?", id).First(&article)
	
	if article.Id > 0 {
		return true
	}
	return false
}

func GetArticleById(id int) *Article {
	article := &Article{}
	db.Model(&Article{}).Where("id = ?", id).Take(article)
	
	db.Model(&article).Related(&article.Tag)
	return article
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Select("count(1)").Count(&count)
	//db.Model(&Article{}).Where(maps).Count(&count)
	return
	
}

/*
Preload就是一个预加载器，它会执行两条 SQL，分别是SELECT * FROM blog_articles;和SELECT * FROM blog_tag WHERE id IN (1,2,3,4);，
那么在查询出结构后，gorm内部处理对应的映射逻辑，将其填充到Article的Tag中，会特别方便，并且避免了循环查询

那么有没有别的办法呢，大致是两种

gorm的Join
循环Related
*/

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	//db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	db.Model(&Article{}).Related(&Tag{})
	db.Model(&Article{}).Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	
	return
}

//func GetArticles() []Article {
//	articles := make([]Article, 128)
//	db.Model(&Article{}).Find(&articles)
//	return articles
//}

func AddArticle(tagId, state int, title, desc, content, createdBy string) {
	db.Model(&Article{}).Create(&Article{
		TagId:     tagId,
		State:     state,
		Title:     title,
		Desc:      desc,
		Content:   content,
		CreatedBy: createdBy,
	})
}

func UpdateArticleById(id int, maps interface{}) bool {
	return db.Model(&Article{}).Where("id = ?", id).Update(maps).Error == nil
}

func DeleteArticleById(id int) bool {
	return db.Where("id = ?", id).Delete(&Article{}).Error == nil
}
