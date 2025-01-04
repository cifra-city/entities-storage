package db

import (
	"context"

	"github.com/cifra-city/entities-storage/internal/data/db/sqlcore"
	"github.com/google/uuid"
)

type PlacesStaff interface {
	Create(ctx context.Context, placeId uuid.UUID, userId uuid.UUID, role string) (sqlcore.PlacesStaff, error)

	Get(ctx context.Context, Id uuid.UUID) (sqlcore.PlacesStaff, error)
	GetByUser(ctx context.Context, placeId uuid.UUID, userId uuid.UUID) (sqlcore.PlacesStaff, error)

	Update(ctx context.Context, Id uuid.UUID, role string) (sqlcore.PlacesStaff, error)

	Delete(ctx context.Context, Id uuid.UUID) error
	DeleteUser(ctx context.Context, placeId uuid.UUID, userId uuid.UUID) error

	ListStaff(ctx context.Context, placeId uuid.UUID) ([]sqlcore.PlacesStaff, error)
}

type placesStaff struct {
	queries *sqlcore.Queries
}

func NewPlacesStaff(queries *sqlcore.Queries) PlacesStaff {
	return &placesStaff{queries: queries}
}

func (d *placesStaff) Create(ctx context.Context, placeId uuid.UUID, userId uuid.UUID, role string) (sqlcore.PlacesStaff, error) {
	return d.queries.CreatePlaceStaff(ctx, sqlcore.CreatePlaceStaffParams{
		ID:      uuid.New(),
		PlaceID: placeId,
		UserID:  userId,
		Role:    role,
	})
}

func (d *placesStaff) Get(ctx context.Context, Id uuid.UUID) (sqlcore.PlacesStaff, error) {
	return d.queries.GetPlaceStaffByID(ctx, Id)
}

func (d *placesStaff) GetByUser(ctx context.Context, placeId uuid.UUID, userId uuid.UUID) (sqlcore.PlacesStaff, error) {
	return d.queries.GetPlaceStaffByPlaceIDAndUserID(ctx, sqlcore.GetPlaceStaffByPlaceIDAndUserIDParams{
		PlaceID: placeId,
		UserID:  userId,
	})
}

func (d *placesStaff) Update(ctx context.Context, Id uuid.UUID, role string) (sqlcore.PlacesStaff, error) {
	return d.queries.UpdatePlaceStaff(ctx, sqlcore.UpdatePlaceStaffParams{
		ID:   Id,
		Role: role,
	})
}

func (d *placesStaff) Delete(ctx context.Context, Id uuid.UUID) error {
	return d.queries.DeletePlaceStaff(ctx, Id)
}

func (d *placesStaff) DeleteUser(ctx context.Context, placeId uuid.UUID, userId uuid.UUID) error {
	return d.queries.DeleteDistributorStaffByDistributorIDAndUserId(ctx, sqlcore.DeleteDistributorStaffByDistributorIDAndUserIdParams{
		DistributorsID: placeId,
		UserID:         userId,
	})
}

func (d *placesStaff) ListStaff(ctx context.Context, placeId uuid.UUID) ([]sqlcore.PlacesStaff, error) {
	return d.queries.ListPlaceStaff(ctx, placeId)
}
