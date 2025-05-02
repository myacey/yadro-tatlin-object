package entity

import (
	"errors"

	"yadro-tatlin-object/internal/models/dto/response"
	"yadro-tatlin-object/internal/models/storage"
)

type Pet struct {
	ASCII       string
	Description string
}

func (c *Pet) MarshalJSON() ([]byte, error) {
	return nil, errors.New("entity.Pet: direct JSON serialization forbidden, use response.Pet")
}

func (c *Pet) ToResponse() *response.Pet {
	return &response.Pet{
		ASCII:       c.ASCII,
		Description: c.Description,
	}
}

func (c *Pet) ToStorage() *storage.Pet {
	return &storage.Pet{
		ASCII:       c.ASCII,
		Description: c.Description,
	}
}
