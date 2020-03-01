package db

import (
	"github.com/jinzhu/gorm"

	"github.com/speeedy/models"
)

func Migrate(db *gorm.DB) {
	objects := []interface{}{
		&models.User{},
		&models.Result{},
	}

	for _, item := range objects {
		if item == nil {
			continue
		}
		db.AutoMigrate(item)
	}

	var user models.User
	if err := db.Where("sender_id = 0").Find(&user).Error; err != nil {
		db.Create(&models.User{
			SenderID:     0,
			UserName:     "system",
			FirstName:    "system",
			LastName:     "system",
			LanguageCode: "go",
			IsBot:        true,
		})
	}
}
