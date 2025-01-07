package models

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Place struct {
	ID          uuid.UUID          `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	ReviewSum   int                `bson:"review_sum,omitempty" json:"review_sum,omitempty"`
	ReviewCount int                `bson:"review_count,omitempty" json:"review_count,omitempty"`
	Type        string             `bson:"type" json:"type"`
	Tags        []string           `bson:"tags" json:"tags"` // Динамический массив тегов
	Schedule    []Schedule         `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Location    GeoPoint           `bson:"location" json:"location"`
	Distributor uuid.UUID          `bson:"distributor" json:"distributor"`
	CityID      uuid.UUID          `bson:"city_id" json:"city_id"`
	CreatedAt   primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt   primitive.DateTime `bson:"updated_at" json:"updated_at"`
}

type GeoPoint struct {
	Type      string  `bson:"type" json:"type"`
	Latitude  float64 `bson:"latitude" json:"latitude"`
	Longitude float64 `bson:"longitude" json:"longitude"`
}

type Schedule struct {
	DayOfWeek time.Weekday `bson:"day_of_week" json:"day_of_week"`
	Open      time.Time    `bson:"open" json:"open"`
	Close     time.Time    `bson:"close" json:"close"`
}
