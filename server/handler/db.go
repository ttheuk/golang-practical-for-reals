package handler

import (
	"source/mix/server/entity"
)

func MigrateDB() {
	DB.Migrate()
}
