package config

import (
	"github.com/masred/mini-blog/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/mini_blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.BlogPost{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
