package db

import (
	"database/sql"

	"github.com/cifra-city/entities-storage/internal/data/db/sqlcore"
)

type Databaser struct {
	Distributors      Distributors
	Places            Places
	PlacesStaff       PlacesStaff
	DistributorsStaff DistributorsStaff
	PlacesTypes       PlacesTypes
	PlacesSchedule    Schedule
}

func NewDBConnection(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func NewDatabaser(url string) (*Databaser, error) {
	db, err := NewDBConnection(url)
	if err != nil {
		return nil, err
	}
	queries := sqlcore.New(db)
	return &Databaser{
		Distributors:      NewDistributors(queries),
		Places:            NewPlaces(queries),
		PlacesStaff:       NewPlacesStaff(queries),
		DistributorsStaff: NewDistributorsStaff(queries),
		PlacesTypes:       NewPlacesTypes(queries),
		PlacesSchedule:    NewSchedule(queries),
	}, nil
}
