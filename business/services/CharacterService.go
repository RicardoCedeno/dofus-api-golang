package services

import (
	// "dofus-api/business/repositories"
	"dofus-api/interfaces/repositories"
	"dofus-api/interfaces/services"
	"dofus-api/models"
)

type CharacterService struct {
	characterRepository repositories.ICharacterRepository
}

func NewCharacterService(repo repositories.ICharacterRepository) services.ICharacterService {
	return &CharacterService{characterRepository: repo}
}

func (s *CharacterService) GetAllCharacters() ([]models.Character, error) {
	return s.characterRepository.GetAllCharacters()
}

func (s *CharacterService) CreateCharacter(character models.Character) error {
	return s.characterRepository.CreateCharacter(character)
}
