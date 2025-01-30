package databases

import (
	"context"
	"log"

	"github.com/nguitarpb/carbowize-test/configs"
	"github.com/nguitarpb/carbowize-test/pkg/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgreSQLDBConnection(cfg *configs.Configs) (*pgxpool.Pool, error) {
	postgresUrl, err := utils.ConnectionUrlBuilder("postgresql", cfg)
	if err != nil {
		return nil, err
	}

	// Initialize connection pool
	poolConfig, err := pgxpool.ParseConfig(postgresUrl)
	if err != nil {
		log.Printf("error parsing config: %s", err.Error())
		return nil, err
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Printf("error connecting to database: %s", err.Error())
		return nil, err
	}

	log.Println("postgreSQL database has been connected 🐘")
	return dbpool, nil
}
