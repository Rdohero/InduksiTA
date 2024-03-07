package initializers

import "InduksiTA/models"

func MigrateDatabase() {
	DB.AutoMigrate(&models.User{})
}
