package config

import (
	"fmt"
	"net/url"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Application struct {
		Port           string
		Name           string
		AllowedOrigins []string
		Location       *time.Location
	}
	PostgreSql struct {
		Driver             string
		Host               string
		Port               string
		Username           string
		Password           string
		Database           string
		DSN                string
		MaxOpenConnections int
		MaxIdleConnections int
	}
	JWT struct {
		PrivateKey string
	}
	Logger struct {
		Formatter logrus.Formatter
	}
}

func Load() *Config {
	cfg := new(Config)

	cfg.application()
	cfg.postgreSql()
	cfg.privateKey()
	cfg.logFormatter()

	return cfg
}

func (cfg *Config) application() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	appName := os.Getenv("APP_NAME")
	port := os.Getenv("PORT")
	timezone := os.Getenv("TIMEZONE")

	if l, err := time.LoadLocation(timezone); err == nil {
		loc = l
	}

	rawAllowedOrigins := strings.Trim(os.Getenv("ALLOWED_ORIGINS"), " ")

	allowedOrigins := make([]string, 0)
	if rawAllowedOrigins == "" {
		allowedOrigins = append(allowedOrigins, "*")
	} else {
		allowedOrigins = strings.Split(rawAllowedOrigins, ",")
	}

	cfg.Application.Port = port
	cfg.Application.Name = appName
	cfg.Application.AllowedOrigins = allowedOrigins
	cfg.Application.Location = loc
}

func (cfg *Config) postgreSql() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	connVal := url.Values{}
	connVal.Add("parseTime", "true")
	connVal.Add("loc", "Asia/Jakarta")

	dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, database)
	dsn := dbConnectionString

	cfg.PostgreSql.Driver = "pgx"
	cfg.PostgreSql.Host = host
	cfg.PostgreSql.Port = port
	cfg.PostgreSql.Username = username
	cfg.PostgreSql.Password = password
	cfg.PostgreSql.Database = database
	cfg.PostgreSql.DSN = dsn

	cfg.PostgreSql.MaxOpenConnections = 10
	cfg.PostgreSql.MaxIdleConnections = 1
}

func (cfg *Config) privateKey() {
	privateKey := os.Getenv("JWT_KEY")
	cfg.JWT.PrivateKey = privateKey
}

func (cfg *Config) logFormatter() {
	formatter := &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			// _, filename := path.Split(f.File)
			filename := fmt.Sprintf("%s:%d", f.File, f.Line)
			return funcname, filename
		},
	}

	cfg.Logger.Formatter = formatter
}
