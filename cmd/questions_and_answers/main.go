package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

//DbConfig environment database config variables
type dbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

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
		SetHosts([]string{fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)}).
		SetAuth(mongoptions.Credential{
			AuthMechanism: "SCRAM-SHA-256",
			Username:      cfg.User,
			Password:      cfg.Password,
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

	repo := questions.NewRepository(mongoClient, cfg.Database, logger)
	svc := questions.NewService(repo, logger)
	eps := questions.MakeEndpoints(svc, logger, middlewares)
	addQuestionEndpoints(r, eps, options)

	level.Info(logger).Log("status", "listening", "port", "8080")
	svr := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: r,
	}
	_ = level.Error(logger).Log(svr.ListenAndServe())
}

func addQuestionEndpoints(rtr *mux.Router, eps questions.Endpoints, options []kitHttp.ServerOption) {
	//All Questions
	rtr.Methods(http.MethodGet).Path("/questions").Handler(transport.GetAllQuestionsHandler(eps.GetAll, options))

	{
		//Question by ID
		path := fmt.Sprintf("/questions/{%s}", transport.QuestionID)
		rtr.Methods(http.MethodGet).Path(path).Handler(transport.GetQuestionHandler(eps.GetQuestion, options))
	}

	//Create Question
	rtr.Methods(http.MethodPost).Path("/questions").Handler(transport.CreateQuestionHandler(eps.CreateQuestion, options))

	{
		//Update question
		path := fmt.Sprintf("/questions/{%s}", transport.QuestionID)
		rtr.Methods(http.MethodPatch).Path(path).Handler(transport.UpdateQuestionHandler(eps.UpdateQuestion, options))
	}

	{
		//Delete question
		path := fmt.Sprintf("/questions/{%s}", transport.QuestionID)
		rtr.Methods(http.MethodDelete).Path(path).Handler(transport.DeleteQuestionHandler(eps.DeleteQuestion, options))
	}
	//Answers related endpoints
	{
		//Add answer to a question
		path := fmt.Sprintf("/questions/{%s}/answers", transport.QuestionID)
		rtr.Methods(http.MethodPut).Path(path).Handler(transport.AddAnswerHandler(eps.CreateAnswer, options))
	}
	{
		//Edit answer of a question
		path := fmt.Sprintf("/questions/{%s}/answers/{%s}", transport.QuestionID, transport.AnswerID)
		rtr.Methods(http.MethodPatch).Path(path).Handler(transport.UpdateAnswerHandler(eps.UpdateAnswer, options))
	}
}

func readConfiguration(logger kitlog.Logger) dbConfig {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return dbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}
}
