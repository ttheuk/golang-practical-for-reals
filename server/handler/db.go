package handler

import (
	"source/mix/server/DB"
)

func MigrateDB() {
	DB.Migrate()
}
