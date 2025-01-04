package db

import (
	"context"

	"github.com/cifra-city/entities-storage/internal/data/db/sqlcore"
)

type PlacesTypes interface {
	Create(ctx context.Context, name string) (sqlcore.PlaceType, error)

	Get(ctx context.Context, id int) (sqlcore.PlaceType, error)
	GetByName(ctx context.Context, name string) (sqlcore.PlaceType, error)

	UpdateName(ctx context.Context, id int, name string) (sqlcore.PlaceType, error)
}

type placesTypes struct {
	queries *sqlcore.Queries
}

func NewPlacesTypes(queries *sqlcore.Queries) PlacesTypes {
	return &placesTypes{queries: queries}
}

func (d *placesTypes) Create(ctx context.Context, name string) (sqlcore.PlaceType, error) {
	return d.queries.CreatePlacesType(ctx, name)
}

func (d *placesTypes) Get(ctx context.Context, id int) (sqlcore.PlaceType, error) {
	return d.queries.GetPlacesTypeByID(ctx, int32(id))
}

func (d *placesTypes) GetByName(ctx context.Context, name string) (sqlcore.PlaceType, error) {
	return d.queries.GetPlacesTypeByName(ctx, name)
}

func (d *placesTypes) UpdateName(ctx context.Context, id int, name string) (sqlcore.PlaceType, error) {
	return d.queries.UpdatePlacesName(ctx, sqlcore.UpdatePlacesNameParams{
		ID:   int32(id),
		Name: name,
	})
}
