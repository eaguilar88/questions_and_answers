package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"bitbucket.org/aveaguilar/questions_and_answers/pkg/entities"
	"bitbucket.org/aveaguilar/questions_and_answers/pkg/questions"
	"bitbucket.org/aveaguilar/questions_and_answers/pkg/transport"
	"github.com/caarlos0/env"
	"github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	mongoptions "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	var logger kitlog.Logger
	{
		logger = kitlog.NewLogfmtLogger(os.Stderr)
		logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)
		logger = kitlog.With(logger, "caller", kitlog.DefaultCaller)
	}

	cfg := readConfiguration(logger)
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	//Database creation
	clientOpts := mongoptions.Client().
		SetConnectTimeout(time.Duration(10) * time.Second).
		SetHosts([]string{fmt.Sprintf("%s:%s", cfg.DB.Host, cfg.DB.Port)}).
		SetAuth(mongoptions.Credential{
			AuthMechanism: "SCRAM-SHA-256",
			Username:      cfg.DB.User,
			Password:      cfg.DB.Password,
		})
	mongoClient, err := mongo.NewClient(clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	err = mongoClient.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(context.Background())

	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		_ = level.Error(logger).Log("msg", "mongo database connection was not established")
		log.Fatal(err)
	}

	var middlewares []endpoint.Middleware
	var options []kitHttp.ServerOption
	r := mux.NewRouter()

	repo := questions.NewRepository(mongoClient, cfg.DB.Database, logger)
	svc := questions.NewService(repo, logger)
	eps := questions.MakeEndpoints(svc, logger, middlewares)
	transport.AddQuestionEndpoints(r, eps, options)

	level.Info(logger).Log("status", "listening", "port", cfg.Port)
	svr := http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Handler: r,
	}
	_ = level.Info(logger).Log(svr.ListenAndServe())
}

func readConfiguration(logger kitlog.Logger) entities.Config {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		level.Error(logger).Log("error", "Error loading .env file")
		os.Exit(1)
	}
	dbConfig := entities.DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}

	return entities.Config{
		Host: os.Getenv("APP_HOST"),
		Port: os.Getenv("APP_PORT"),
		DB:   dbConfig,
	}
}
