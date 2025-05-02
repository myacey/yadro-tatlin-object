package service

import (
	"context"
	"errors"

	"yadro-tatlin-object/internal/models/dto/request"
	"yadro-tatlin-object/internal/models/entity"
	"yadro-tatlin-object/internal/pkg/web/apperror"
	"yadro-tatlin-object/internal/repository"
)

type PetRepo interface {
	GetPet(ctx context.Context) (*entity.Pet, error)
	UploadPet(ctx context.Context, pet *entity.Pet) error
}

type PetService struct {
	repo PetRepo
}

func NewPet(repo PetRepo) *PetService {
	return &PetService{repo: repo}
}

func (s *PetService) GetPet(ctx context.Context) (*entity.Pet, error) {
	pet, err := s.repo.GetPet(ctx)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrPetNotFound):
			return nil, apperror.NewNoContent("no pet uploaded yet")
		default:
			return nil, apperror.NewInternal("failed to get pet", err)
		}
	}

	return pet, nil
}

func (s *PetService) UploadPet(ctx context.Context, req *request.UploadPet) error {
	pet := &entity.Pet{
		ASCII:       req.ASCII,
		Description: req.Description,
	}
	err := s.repo.UploadPet(ctx, pet)
	if err != nil {
		return apperror.NewInternal("failed to upload pet", err)
	}

	return nil
}
