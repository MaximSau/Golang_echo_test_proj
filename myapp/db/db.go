package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Db init
var Db, Err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
