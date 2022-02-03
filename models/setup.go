package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		// panic ให้หยุดการทำงานของระบบแล้วขึ้น log มา
		panic("Failed to connect to database!")
	}

	// สร้าง table ขึ้นใน database
	database.AutoMigrate(&Todo{})

	DB = database
}
