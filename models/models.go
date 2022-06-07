package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type BaseModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

var DB *gorm.DB

func InitDB() (db *gorm.DB, err error) {
	db, err = gorm.Open("sqlite3", "./database/weblog.db")
	if err == nil {
		DB = db
		//db.LogMode(true)
		//db.AutoMigrate(&Page{}, &Post{}, &Tag{}, &PostTag{}, &User{}, &Comment{}, &Subscriber{}, &Link{}, &SmmsFile{})
		//db.Model(&PostTag{}).AddUniqueIndex("uk_post_tag", "post_id", "tag_id")
		db.AutoMigrate()
		return
	}
	return
}
