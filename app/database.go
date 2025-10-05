package app

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// データベース設定
type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dhname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", c.Host, c.User, c.Password, c.DBName, c.Port)
}

func NewDB(c *Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(c.DSN()), &gorm.Config{})
}

// マイグレーションを実行する
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &Auth{}, &Comment{}, &Post{})
}
