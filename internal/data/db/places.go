package db

import (
	"context"

	"github.com/cifra-city/entities-storage/internal/data/db/sqlcore"
	"github.com/google/uuid"
)

type Places interface {
	Create(
		ctx context.Context,
		name string,
		distributorId uuid.UUID,
		streetId uuid.UUID,
		houseNumber string,
		location string,
	) (sqlcore.Place, error)

	Get(ctx context.Context, id uuid.UUID) (sqlcore.Place, error)

	UpdateName(ctx context.Context, id uuid.UUID, name string) (sqlcore.Place, error)
	UpdateLocation(ctx context.Context, id uuid.UUID, location string, houseNumber string) (sqlcore.Place, error)
	UpdateDistributor(ctx context.Context, id uuid.UUID, distributorId uuid.UUID) (sqlcore.Place, error)
	UpdateType(ctx context.Context, id uuid.UUID, typeId int) (sqlcore.Place, error)
	UpdatePlaceScore(ctx context.Context, id uuid.UUID, addScore int) (sqlcore.Place, error)

	Delete(ctx context.Context, id uuid.UUID) error

	List(ctx context.Context) ([]sqlcore.Place, error)
	ListByStreet(ctx context.Context, streetId uuid.UUID) ([]sqlcore.Place, error)
	ListByDistributor(ctx context.Context, distributorId uuid.UUID) ([]sqlcore.Place, error)
	ListByType(ctx context.Context, typeId int) ([]sqlcore.Place, error)
	ListByTypeAndStreet(ctx context.Context, typeId int, streetId uuid.UUID) ([]sqlcore.Place, error)
}

type places struct {
	queries *sqlcore.Queries
}

func NewPlaces(queries *sqlcore.Queries) Places {
	return &places{queries: queries}
}

func (p *places) Create(
	ctx context.Context,
	name string,
	distributorId uuid.UUID,
	streetId uuid.UUID,
	houseNumber string,
	location string,
) (sqlcore.Place, error) {
	return p.queries.CreatePlace(ctx, sqlcore.CreatePlaceParams{
		Name:          name,
		DistributorID: distributorId,
		StreetID:      streetId,
		HouseNumber:   houseNumber,
		Location:      location,
	})
}

func (p *places) Get(ctx context.Context, id uuid.UUID) (sqlcore.Place, error) {
	return p.queries.GetPlaceByID(ctx, id)
}

func (p *places) UpdateName(ctx context.Context, id uuid.UUID, name string) (sqlcore.Place, error) {
	return p.queries.UpdatePlaceName(ctx, sqlcore.UpdatePlaceNameParams{
		ID:   id,
		Name: name,
	})
}

func (p *places) UpdateLocation(ctx context.Context, id uuid.UUID, location string, houseNumber string) (sqlcore.Place, error) {
	return p.queries.UpdatePlaceLocation(ctx, sqlcore.UpdatePlaceLocationParams{
		ID:          id,
		Location:    location,
		HouseNumber: houseNumber,
	})
}

func (p *places) UpdateDistributor(ctx context.Context, id uuid.UUID, distributorId uuid.UUID) (sqlcore.Place, error) {
	return p.queries.UpdatePlaceDistributor(ctx, sqlcore.UpdatePlaceDistributorParams{
		ID:            id,
		DistributorID: distributorId,
	})
}

func (p *places) UpdateType(ctx context.Context, id uuid.UUID, typeId int) (sqlcore.Place, error) {
	return p.queries.UpdatePlaceType(ctx, sqlcore.UpdatePlaceTypeParams{
		ID:   id,
		Type: int32(typeId),
	})
}

func (p *places) UpdatePlaceScore(ctx context.Context, id uuid.UUID, addScore int) (sqlcore.Place, error) {
	return p.queries.UpdatePlaceScore(ctx, sqlcore.UpdatePlaceScoreParams{
		ID:         id,
		TotalScore: int32(addScore),
	})
}

func (p *places) Delete(ctx context.Context, id uuid.UUID) error {
	return p.queries.DeletePlace(ctx, id)
}

func (p *places) List(ctx context.Context) ([]sqlcore.Place, error) {
	return p.queries.ListPlaces(ctx)
}

func (p *places) ListByDistributor(ctx context.Context, distributorId uuid.UUID) ([]sqlcore.Place, error) {
	return p.queries.ListPlacesByDistributor(ctx, distributorId)
}

func (p *places) ListByStreet(ctx context.Context, streetId uuid.UUID) ([]sqlcore.Place, error) {
	return p.queries.ListPlacesByStreet(ctx, streetId)
}

func (p *places) ListByType(ctx context.Context, typeId int) ([]sqlcore.Place, error) {
	return p.queries.ListPlacesByType(ctx, int32(typeId))
}

func (p *places) ListByTypeAndStreet(ctx context.Context, typeId int, streetId uuid.UUID) ([]sqlcore.Place, error) {
	return p.queries.ListPlacesByStreetAndType(ctx, sqlcore.ListPlacesByStreetAndTypeParams{
		StreetID: streetId,
		Type:     int32(typeId),
	})
}
