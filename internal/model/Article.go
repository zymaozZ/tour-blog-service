package model

type Article struct {
	*Model
	ID            uint   `gorm:"primaryKey;column:id" json:"-"`
	Title         string `gorm:"column:title" json:"title"`                   // 文章标题
	Desc          string `gorm:"column:desc" json:"desc"`                     // 文章简述
	CoverImageURL string `gorm:"column:cover_image_url" json:"coverImageUrl"` // 封面图片地址
	Content       string `gorm:"column:content" json:"content"`               // 文章内容
	State         uint8  `gorm:"column:state" json:"state"`                   // 状态 0为禁用、1为启用
}
