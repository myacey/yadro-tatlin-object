package repository

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"sync"

	"yadro-tatlin-object/internal/models/entity"
	"yadro-tatlin-object/internal/models/storage"
)

var ErrPetNotFound = errors.New("pet not found")

type Config struct {
	DataDir  string `mapstructure:"data_dir"`
	FileName string `mapstructure:"filename"`
}

type PetRepo struct {
	dataDir  string
	fileName string
	mu       sync.RWMutex
}

func NewPet(cfg Config) *PetRepo {
	return &PetRepo{dataDir: cfg.DataDir, fileName: cfg.FileName}
}

func (r *PetRepo) savePetToFile(pet *storage.Pet) error {
	filePath := r.dataDir + r.fileName

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(pet)
}

func (r *PetRepo) loadPetFromFile() (*storage.Pet, error) {
	file, err := os.Open(r.dataDir + r.fileName)
	if err != nil {
		return nil, ErrPetNotFound
	}
	defer file.Close()

	var pet storage.Pet
	if err = json.NewDecoder(file).Decode(&pet); err != nil {
		return nil, err
	}

	return &pet, nil
}

func (r *PetRepo) GetPet(ctx context.Context) (*entity.Pet, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	pet, err := r.loadPetFromFile()
	if err != nil {
		return nil, err
	}
	return &entity.Pet{
		ASCII:       pet.ASCII,
		Description: pet.Description,
	}, nil
}

func (r *PetRepo) UploadPet(ctx context.Context, pet *entity.Pet) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.savePetToFile(pet.ToStorage())
}
