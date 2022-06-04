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

var Db *gorm.DB

func InitDB() (db *gorm.DB, err error) {
	db, err = gorm.Open("sqlite3", "./database/weblog.db")
	return
}
