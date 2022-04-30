package model

import (
	"blog-service/global"
	"blog-service/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID         uint   `gorm:"primaryKey;column:id" json:"id"`
	CreatedOn  uint   `gorm:"column:created_on" json:"createdOn"`   // 创建时间
	CreatedBy  string `gorm:"column:created_by" json:"createdBy"`   // 创建人
	ModifiedOn uint   `gorm:"column:modified_on" json:"modifiedOn"` // 修改时间
	ModifiedBy string `gorm:"column:modified_by" json:"modifiedBy"` // 修改人
	DeletedOn  uint   `gorm:"column:deleted_on" json:"deletedOn"`   // 删除时间
	IsDel      uint8  `gorm:"column:is_del" json:"isDel"`           // 是否删除 0为未删除、1为已删除
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}
