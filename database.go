package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(c *Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(c.DSN()), &gorm.Config{})
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &Auth{}, &Comment{}, &Post{})
}
