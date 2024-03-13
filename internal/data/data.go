package data

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func init() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

type Data struct {
	Db    *gorm.DB
	redis *redis.Client
}

// NewMysql
func NewMysql() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: viper.Get("database.mysql.dsn").(string),
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
		Addr:     viper.Get("database.redis.addr").(string),
		Password: viper.Get("database.redis.password").(string), // no password set
		DB:       viper.Get("database.redis.db").(int),          // use default DB
	})
	return client
}

func NewData(db *gorm.DB, redisCli *redis.Client) *Data {
	return &Data{Db: db, redis: redisCli}
}
