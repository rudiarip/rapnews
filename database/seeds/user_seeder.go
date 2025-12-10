package seeds

import (
	"rapnews/internal/core/domain/model"
	"rapnews/lib/conv"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	bytes, err := conv.HashPassword("admin123")
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating Password Hash")
	}

	admin := model.User{
		Name:     "admin",
		Email:    "admin@gmail.com",
		Password: string(bytes),
	}

	if err := db.FirstOrCreate(&admin, model.User{Email: "admin@gmail.com"}).Error; err != nil {
		log.Fatal().Err(err).Msg("Error seeding admin user")
	} else {
		log.Info().Msg("Admin user created")
	}
}
