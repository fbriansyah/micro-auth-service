package main

import (
	grpcclient "github.com/fbriansyah/micro-auth-service/internal/adapter/client/grpc"
	"github.com/fbriansyah/micro-auth-service/internal/adapter/postgresdb"
	grpcserver "github.com/fbriansyah/micro-auth-service/internal/adapter/server/grpc"
	"github.com/fbriansyah/micro-auth-service/internal/application"
	"github.com/fbriansyah/micro-auth-service/util"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	var opts []grpc.DialOption

	databaseAdapter := postgresdb.NewDatabaseAdapter(sqlDB)

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(config.SessionServerAddress, opts...)
	if err != nil {
		log.Fatal().Msgf("cannot connect to session service: %s", err.Error())
	}
	defer conn.Close()

	sessionClient := grpcclient.NewSessionAdapterClient(conn)

	authService := application.NewAuthService(databaseAdapter, sessionClient)

	serverAdapter := grpcserver.NewGrpcServerAdapter(authService, config.GrpcPort)
	serverAdapter.Run()
}
