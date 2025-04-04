package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/jpmoraess/appointment-api/db/sqlc"
	_ "github.com/jpmoraess/appointment-api/docs"
	"github.com/jpmoraess/appointment-api/gapi"
	"github.com/jpmoraess/appointment-api/internal/handlers"
	"github.com/jpmoraess/appointment-api/internal/router"
	"github.com/jpmoraess/appointment-api/internal/usecases"
	"github.com/jpmoraess/appointment-api/pb"
	"github.com/jpmoraess/appointment-api/pkg/metrics"
	"github.com/jpmoraess/appointment-api/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// @title			Appointment System
// @version		1.0
// @description	Appointment System Documentation
// @termsOfService	https://example.com/terms
// @host			localhost:8080
func main() {
	cfg, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dbConfig, err := pgxpool.ParseConfig(cfg.DbSource)
	if err != nil {
		log.Fatal("cannot parse connection:", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatal("cannot create connection to database: ", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	defer pool.Close()

	store := db.NewStore(pool)

	// uc's
	register := usecases.NewRegister(store)
	_ = register

	//runHttpServer(cfg, register)

	runGrpcServer(store, cfg, register)
}

func runHttpServer(cfg utils.Config, register *usecases.Register) {
	// fiber
	app := fiber.New(fiber.Config{
		AppName: "Appointment System",
	})

	// middlewares
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(metrics.New())

	// handlers
	registerHandler := handlers.NewRegisterHandler(register)

	// routes
	router := router.NewRouter(app, registerHandler)
	router.SetupRoutes()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Printf("Server running on port: %s", cfg.HttpServerAddress)
		if err := app.Listen(cfg.HttpServerAddress); err != nil {
			log.Fatalf("Error while starting HTTP server: %v", err)
		}
	}()

	<-quit
	log.Println("Shutting down the HTTP server...")

	timeout := 5 * time.Second
	if err := app.ShutdownWithTimeout(timeout); err != nil {
		log.Fatalf("Error when shutting down the HTTP server: %v", err)
	}
	log.Println("Server shut down successfully...")
}

func runGrpcServer(store db.Store, cfg utils.Config, register *usecases.Register) {
	server := gapi.NewServer(cfg, store, register)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", cfg.GrpcServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server at %s", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("cannot start gRPC server")
	}
}
