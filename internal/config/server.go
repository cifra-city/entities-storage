package config

import (
	"github.com/cifra-city/cifra-rabbit"
	"github.com/cifra-city/entities-storage/internal/data/nosql"
	"github.com/cifra-city/entities-storage/internal/data/sql"
	"github.com/cifra-city/tokens"
	"github.com/sirupsen/logrus"
)

const (
	SERVER = "server"
)

type Service struct {
	Config       *Config
	SqlDB        *sql.Repo
	MongoDB      *nosql.Repository
	Logger       *logrus.Logger
	TokenManager *tokens.TokenManager
	Broker       *cifra_rabbit.Broker
}

func NewServer(cfg *Config) (*Service, error) {
	logger := SetupLogger(cfg.Logging.Level, cfg.Logging.Format)
	queries, err := sql.NewRepoSQl(cfg.Database.URL)
	if err != nil {
		return nil, err
	}
	TokenManager := tokens.NewTokenManager(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB, logger, cfg.JWT.AccessToken.TokenLifetime)
	broker, err := cifra_rabbit.NewBroker(cfg.Rabbit.URL, cfg.Rabbit.Exchange)
	if err != nil {
		return nil, err
	}
	mogoRepo, err := nosql.NewRepository(cfg.MongoDB.URL, cfg.MongoDB.database)

	return &Service{
		Config:       cfg,
		SqlDB:        queries,
		MongoDB:      mogoRepo,
		Logger:       logger,
		TokenManager: TokenManager,
		Broker:       broker,
	}, nil
}
