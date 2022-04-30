package model

type ArticleTag struct {
	*Model
	ArticleID int  `gorm:"column:article_id" json:"articleId"` // 文章ID
	TagID     uint `gorm:"column:tag_id" json:"tagId"`         // 标签ID
}
