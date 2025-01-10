package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload" // for development
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/umerthow/crud-api/config"
	"github.com/umerthow/crud-api/response"
	"github.com/umerthow/crud-api/server"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/v2/pgxv4"
)

var (
	cfg          *config.Config
	location     *time.Location
	indexMessage string = "Application is running properly"
)

func init() {
	cfg = config.Load()
}

func main() {
	logger := logrus.New()
	logger.SetFormatter(cfg.Logger.Formatter)

	db, err := apmsql.Open(cfg.PostgreSql.Driver, cfg.PostgreSql.DSN)
	if err != nil {
		logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		logger.Fatal(err)
	} else {
		logger.Println("Database Successfully Connected")
	}

	router := mux.NewRouter()
	router.HandleFunc("/user", index)

	handler := cors.New(cors.Options{
		AllowedOrigins:   cfg.Application.AllowedOrigins,
		AllowedMethods:   []string{http.MethodPost, http.MethodGet, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "X-Requested-With", "Authorization"},
		AllowCredentials: true,
	}).Handler(router)

	// initiate server
	srv := server.NewServer(logger, handler, cfg.Application.Port)
	srv.Start()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm

	// closing service for a gracefull shutdown.
	srv.Close()
}

func index(w http.ResponseWriter, r *http.Request) {
	resp := response.NewSuccessResponse(nil, response.StatOK, indexMessage)
	response.JSON(w, resp)
}
