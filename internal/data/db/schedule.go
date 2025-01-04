package db

import (
	"context"
	"time"

	"github.com/cifra-city/entities-storage/internal/data/db/sqlcore"
	"github.com/google/uuid"
)

type Schedule interface {
	Create(ctx context.Context, placeID uuid.UUID, dayWeek int, openTime time.Time, closeTime time.Time) (sqlcore.PlaceSchedule, error)

	Get(ctx context.Context, id uuid.UUID) (sqlcore.PlaceSchedule, error)
	GetByPlaceIdAndDay(ctx context.Context, placeID uuid.UUID, dayWeek int) (sqlcore.PlaceSchedule, error)

	ListByPlaceId(ctx context.Context, placeID uuid.UUID) ([]sqlcore.PlaceSchedule, error)
	ListByDay(ctx context.Context, dayWeek int) ([]sqlcore.PlaceSchedule, error)

	Update(ctx context.Context, id uuid.UUID, dayWeek int, openTime time.Time, closeTime time.Time) (sqlcore.PlaceSchedule, error)
	UpdateByPlaceId(ctx context.Context, placeID uuid.UUID, dayWeek int, openTime time.Time, closeTime time.Time) (sqlcore.PlaceSchedule, error)

	Delete(ctx context.Context, id uuid.UUID) error
}

type schedule struct {
	queries *sqlcore.Queries
}

func NewSchedule(queries *sqlcore.Queries) Schedule {
	return &schedule{queries: queries}
}

func (s *schedule) Create(ctx context.Context, placeID uuid.UUID, dayWeek int, openTime time.Time, closeTime time.Time) (sqlcore.PlaceSchedule, error) {
	return s.queries.CreateSchedule(ctx, sqlcore.CreateScheduleParams{
		PlaceID:   placeID,
		DayOfWeek: int32(dayWeek),
		OpenTime:  openTime,
		CloseTime: closeTime,
	})
}

func (s *schedule) Get(ctx context.Context, id uuid.UUID) (sqlcore.PlaceSchedule, error) {
	return s.queries.GetScheduleByID(ctx, id)
}

func (s *schedule) GetByPlaceIdAndDay(ctx context.Context, placeID uuid.UUID, dayWeek int) (sqlcore.PlaceSchedule, error) {
	return s.queries.GetScheduleByPlaceIDAndDay(ctx, sqlcore.GetScheduleByPlaceIDAndDayParams{
		PlaceID:   placeID,
		DayOfWeek: int32(dayWeek),
	})
}

func (s *schedule) ListByPlaceId(ctx context.Context, placeID uuid.UUID) ([]sqlcore.PlaceSchedule, error) {
	return s.queries.ListScheduleByPlaceID(ctx, placeID)
}

func (s *schedule) ListByDay(ctx context.Context, dayWeek int) ([]sqlcore.PlaceSchedule, error) {
	return s.queries.ListScheduleByDay(ctx, int32(dayWeek))
}

func (s *schedule) Update(ctx context.Context, Id uuid.UUID, dayWeek int, openTime time.Time, closeTime time.Time) (sqlcore.PlaceSchedule, error) {
	return s.queries.UpdateSchedule(ctx, sqlcore.UpdateScheduleParams{
		ID:        Id,
		DayOfWeek: int32(dayWeek),
		OpenTime:  openTime,
		CloseTime: closeTime,
	})
}

func (s *schedule) UpdateByPlaceId(ctx context.Context, placeID uuid.UUID, dayWeek int, openTime time.Time, closeTime time.Time) (sqlcore.PlaceSchedule, error) {
	return s.queries.UpdateScheduleByPlaceId(ctx, sqlcore.UpdateScheduleByPlaceIdParams{
		PlaceID:   placeID,
		DayOfWeek: int32(dayWeek),
		OpenTime:  openTime,
		CloseTime: closeTime,
	})
}

func (s *schedule) Delete(ctx context.Context, id uuid.UUID) error {
	return s.queries.DeleteSchedule(ctx, id)
}
