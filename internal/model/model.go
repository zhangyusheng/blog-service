package model

import (
	"fmt"
	"github.com/zhangyusheng/blog-service/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	dsn := fmt.Sprintf(s, databaseSetting.UserName, databaseSetting.Password, databaseSetting.Host, databaseSetting.DBName, databaseSetting.Charset, databaseSetting.ParseTime)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(any(err))
	}
	sql, err := db.DB()
	sql.SetMaxIdleConns(10)
	sql.SetMaxOpenConns(100)
	sql.SetConnMaxIdleTime(time.Hour)
	return db, nil
}
