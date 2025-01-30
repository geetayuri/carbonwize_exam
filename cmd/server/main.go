package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nguitarpb/carbowize-test/configs"
	"github.com/nguitarpb/carbowize-test/modules/servers"
	databases "github.com/nguitarpb/carbowize-test/pkg/databases"

	"github.com/nguitarpb/carbowize-test/virtualfuncforfifth"
)

func main() {
	activity := "transportation"
	input := 75
	var efc float32 = 0.41

	fifthOutput := virtualfuncforfifth.CalculateEmission(&activity, &input, &efc)
	fmt.Println(*fifthOutput)

	// Load dotenv config
	if err := godotenv.Load("./.env"); err != nil {
		panic(err.Error())
	}
	cfg := new(configs.Configs)

	// Fiber configs
	cfg.App.Host = os.Getenv("FIBER_HOST")
	cfg.App.Port = os.Getenv("FIBER_PORT")

	// Database Configs
	cfg.PostgreSQL.Host = os.Getenv("POSTGRESDB_HOST")
	cfg.PostgreSQL.Port = os.Getenv("POSTGRESDB_PORT")
	cfg.PostgreSQL.Protocol = os.Getenv("POSTGRESDB_PROTOCOL")
	cfg.PostgreSQL.Username = os.Getenv("POSTGRESDB_USERNAME")
	cfg.PostgreSQL.Password = os.Getenv("POSTGRESDB_PASSWORD")
	cfg.PostgreSQL.Database = os.Getenv("POSTGRESDB_DATABASE")
	cfg.PostgreSQL.SSLMode = os.Getenv("POSTGRESDB_SSL_MODE")

	// New Database
	postgresdb, err := databases.NewPostgreSQLDBConnection(cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer postgresdb.Close()

	s := servers.NewServer(cfg, postgresdb)
	s.Start()
}
