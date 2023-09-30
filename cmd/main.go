package main

import (
	grpcclient "github.com/fbriansyah/micro-auth-service/internal/adapter/client/grpc"
	"github.com/fbriansyah/micro-auth-service/internal/adapter/postgresdb"
	grpcserver "github.com/fbriansyah/micro-auth-service/internal/adapter/server/grpc"
	"github.com/fbriansyah/micro-auth-service/internal/application"
	"github.com/fbriansyah/micro-auth-service/util"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal().Msgf("cannot load config: %s", err.Error())
	}

	sqlDB := connectToDB(config.DBDriver, config.DBSource)
	if sqlDB == nil {
		log.Fatal().Msgf("cannot connect to db: %s", err.Error())
	}

	runDBMigration(config.MigrationURL, config.DBSource)

	databaseAdapter := postgresdb.NewDatabaseAdapter(sqlDB)
	conn, err := grpc.Dial(config.SessionServerAddress, nil)
	if err != nil {
		log.Fatal().Msgf("cannot connect to session service: %s", err.Error())
	}
	defer conn.Close()

	sessionClient := grpcclient.NewSessionAdapterClient(conn)

	authService := application.NewAuthService(databaseAdapter, sessionClient)

	grpcserver.NewGrpcServerAdapter(authService, config.GrpcPort)
}
