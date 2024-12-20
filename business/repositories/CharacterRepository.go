package repositories

import (
	"dofus-api/interfaces/repositories"
	"dofus-api/models"

	"gorm.io/gorm"
)

type CharacterRepository struct {
	DB *gorm.DB
}

func NewCharacterRepository(db *gorm.DB) repositories.ICharacterRepository {
	return &CharacterRepository{DB: db}
}

func (r *CharacterRepository) GetAllCharacters() ([]models.Character, error) {
	var characters []models.Character
	if err := r.DB.Find(&characters).Error; err != nil {
		return nil, err
	}
	return characters, nil
}

func (r *CharacterRepository) CreateCharacter(character models.Character) error {
	if err := r.DB.Create(character).Error; err != nil {
		return err
	}
	return nil
}
