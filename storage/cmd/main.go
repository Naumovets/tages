package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"

	"github.com/Naumovets/tages/config"
	controller "github.com/Naumovets/tages/internal/controller/grpc"
	"github.com/Naumovets/tages/internal/db/postgres"
	"github.com/Naumovets/tages/internal/repository"
	tages "github.com/Naumovets/tages/pkg/proto/storage"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const (
	grpcAddress = "0.0.0.0:50051"
)

func main() {
	cfg := config.New()

	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.USER,
		cfg.DB.PASSWORD,
		cfg.DB.DB_HOST,
		cfg.DB.DB_PORT,
		cfg.DB.DB_NAME,
	)
	m, err := migrate.New("file://tools/migrations", databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	version, dirty, err := m.Version()

	if err != nil {
		log.Fatal(err)
	}

	slog.Info(fmt.Sprintf("Applied migration: %d, Dirty: %t\n", version, dirty))

	conn, err := postgres.NewConn(cfg)

	if err != nil {
		log.Fatalf("err: %s\n", err)
		os.Exit(1)
	}

	rep := repository.NewRepository(conn)

	if err != nil {
		log.Fatalf("err: %s\n", err)
		os.Exit(2)
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)

	reflection.Register(grpcServer)

	tages.RegisterStorageServer(grpcServer, controller.NewServerStorage(cfg, rep))

	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("gRPC server listening")

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
