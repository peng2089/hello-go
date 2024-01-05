package data

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Data struct {
	db *gorm.DB
}

func NewMysql() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(localhost:3306)/hello-world?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tbl_", // 表名前缀
			SingularTable: true,   // 使用单数表名
		}})
	if err != nil {
		fmt.Printf("Err: %+v\n", err)
		return nil, err
	}
	return db, nil
}

func NewData(db *gorm.DB) *Data {
	return &Data{db: db}
}
