package model

import (
	"fmt"
	"github.com/cnodin/blog-service/global"
	"github.com/cnodin/blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
)

type Model struct {
	ID uint32	`gorm:"primary_key" json:"id"`
	CreatedBy	string `json:"created_by"`
	ModifiedBy	string	`json:"modified_by"`
	CreatedOn	uint32	`json:"created_on"`
	ModifiedOn	uint32	`json:"modified_on"`
	DeletedOn	uint32	`json:"deleted_on"`
	IsDel		uint8	`json:"is_del"`
}

func NewDBEngine(dbSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(dbSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			dbSetting.UserName,
			dbSetting.Password,
			dbSetting.Host,
			dbSetting.DBName,
			dbSetting.Charset,
			dbSetting.ParseTime,
		))
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(dbSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(dbSetting.MaxOpenConns)

	return db, nil
}
