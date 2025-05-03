package main

import "fmt"

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
