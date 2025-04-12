package migrations

import (
	"Cars/internal/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func GetMigrations(db *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20250412_create_tables",
			Migrate: func(tx *gorm.DB) error {
				// Сначала создаем таблицу ролей
				if err := tx.AutoMigrate(&models.Role{}); err != nil {
					return err
				}

				// Затем создаем таблицу пользователей с внешним ключом
				if err := tx.AutoMigrate(&models.User{}, &models.Car{}); err != nil {
					return err
				}

				// Создаем таблицу для refresh токенов
				return tx.AutoMigrate(&models.RefreshToken{})
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Migrator().DropTable("refresh_tokens"); err != nil {
					return err
				}
				if err := tx.Migrator().DropTable("users", "cars"); err != nil {
					return err
				}
				return tx.Migrator().DropTable("role")
			},
		},
	})
}
