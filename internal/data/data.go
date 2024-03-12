package data

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Data struct {
	Db    *gorm.DB
	redis *redis.Client
}

// NewMysql
func NewMysql() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(mysql:3306)/hello-go?charset=utf8mb4&parseTime=True&loc=Local",
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

// NewRedis
func NewRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

func NewData(db *gorm.DB, redisCli *redis.Client) *Data {
	return &Data{Db: db, redis: redisCli}
}
