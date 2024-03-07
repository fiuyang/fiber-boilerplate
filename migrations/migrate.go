package migrations

import (
	"fiber-boilerplate/model"

	"gorm.io/gorm"
)

func CheckTableNotExists(db *gorm.DB, modelInstance interface{}) {
	if !db.Migrator().HasTable(modelInstance) {
		db.AutoMigrate(modelInstance)
	}
}

func AutoMigrate(db *gorm.DB) {
	CheckTableNotExists(db, &model.Note{})
}