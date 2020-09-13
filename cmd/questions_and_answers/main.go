package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"bitbucket.org/aveaguilar/questions_and_answers/pkg/answers"
	"bitbucket.org/aveaguilar/questions_and_answers/pkg/questions"
	"github.com/caarlos0/env"
	"github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kithttp "github.com/go-kit/kit/transport/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

//DbConfig environment database config variables
type dbConfig struct {
	Driver   string
	Host     string
	Port     int
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

	cfg := readConfiguration()
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	db, err := NewDb(cfg)
	if err != nil {
		level.Error(logger).Log("error", err)
	}

	var middlewares []endpoint.Middleware
	var options []kithttp.ServerOption
	r := mux.NewRouter()

	repo := questions.NewRepository(db, logger)
	svc := questions.NewService(repo, logger)
	eps := questions.MakeEndpoints(svc, logger, middlewares)
	addQuestionEndpoints(r, eps, options)

	answerRepo := answers.NewRepository(db, logger)
	answerSvc := answers.NewService(answerRepo, logger)
	answerEps := answers.MakeEndpoints(answerSvc, logger, middlewares)

	addAnswerEndpoints(r, answerEps, options)
	level.Info(logger).Log("status", "listening", "port", "8080")
	svr := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: r,
	}
	level.Error(logger).Log(svr.ListenAndServe())
}

func addQuestionEndpoints(rtr *mux.Router, eps questions.Endpoints, options []kithttp.ServerOption) {
	rtr.Methods(http.MethodGet).Path("/questions/all").Handler(questions.GetAllQuestionsHandler(eps.GetAll, options))

	{
		path := fmt.Sprintf("/questions/{%s}", params.DomainID)
		rtr.Methods(http.MethodGet).Path(path).Handler(questions.GetAllQuestionsHandler(eps.GetAll, options))
	}
	rtr.Methods(http.MethodPost).Path("/questions/create").Handler(questions.CreateQuestionHandler(eps.CreateQuestion, options))
}

func addAnswerEndpoints(rtr *mux.Router, eps answers.Endpoints, options []kithttp.ServerOption) {
	rtr.Methods(http.MethodGet).Path("/palindrome").Handler(answers.GetIsPalHandler(eps.GetIsPalindrome, options))
	rtr.Methods(http.MethodGet).Path("/reverse").Handler(answers.GetReverseHandler(eps.GetReverse, options))
}

//NewDb _
func NewDb(cfg dbConfig) (*sql.DB, error) {
	database, err := sql.Open(cfg.Driver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database))
	if err != nil {
		return nil, err
	}

	return database, nil
}

func readConfiguration() dbConfig {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return dbConfig{
		Driver:   os.Getenv("DB_DRIVER"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}
}
