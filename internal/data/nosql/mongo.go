package nosql

import (
	"fmt"

	"github.com/cifra-city/entities-storage/internal/data/nosql/repositories/repositories"
)

type Repository struct {
	Places repositories.Places
}

func NewRepository(uri, dbName string) (*Repository, error) {
	placesRepo, err := repositories.NewPlaces(uri, dbName, "places")
	if err != nil {
		return nil, fmt.Errorf("failed to initialize places repository: %w", err)
	}

	return &Repository{
		Places: placesRepo,
	}, nil
}
