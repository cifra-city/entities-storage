package sql

import (
	"database/sql"

	"github.com/cifra-city/entities-storage/internal/data/sql/repositories"
	"github.com/cifra-city/entities-storage/internal/data/sql/repositories/sqlcore"
)

type Repo struct {
	Distributors      repositories.Distributors
	DistributorsStaff repositories.DistributorsStaff
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

func NewRepoSQL(url string) (*Repo, error) {
	db, err := NewDBConnection(url)
	if err != nil {
		return nil, err
	}
	queries := sqlcore.New(db)
	return &Repo{
		Distributors:      repositories.NewDistributors(queries),
		DistributorsStaff: repositories.NewDistributorsStaff(queries),
	}, nil
}
