package main

import (
	"Cars/internal/services"
)

func main() {
	services.ConnectDatabase()
	services.RollbackMigrations()
}
