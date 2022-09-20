package database

import (
	url_service "github.com/cingozr/go-tiny-url/internal/tinyurl"
	"github.com/cingozr/go-tiny-url/internal/user"
	"github.com/jinzhu/gorm"
)

func MigrateDb(db *gorm.DB) error {
	if result := db.AutoMigrate(&url_service.TinyUrl{}, &user.User{}); result.Error != nil {
		return result.Error
	}
	return nil
}
