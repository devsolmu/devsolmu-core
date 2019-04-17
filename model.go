package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Project is main content
type Project struct {
	gorm.Model
	Title string `gorm:"unique" json:"title"`
}

// DBMigrate is automatically migate database
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Project{})
	return db
}
