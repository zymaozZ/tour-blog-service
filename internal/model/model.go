package model

type Model struct {
	ID         uint   `gorm:"primaryKey;column:id" json:"id"`
	CreatedOn  uint   `gorm:"column:created_on" json:"createdOn"`   // 创建时间
	CreatedBy  string `gorm:"column:created_by" json:"createdBy"`   // 创建人
	ModifiedOn uint   `gorm:"column:modified_on" json:"modifiedOn"` // 修改时间
	ModifiedBy string `gorm:"column:modified_by" json:"modifiedBy"` // 修改人
	DeletedOn  uint   `gorm:"column:deleted_on" json:"deletedOn"`   // 删除时间
	IsDel      uint8  `gorm:"column:is_del" json:"isDel"`           // 是否删除 0为未删除、1为已删除
}

//ArticleID  int    `gorm:"column:article_id" json:"articleId"`   // 文章ID
//TagID      uint   `gorm:"column:tag_id" json:"tagId"`           // 标签ID
