package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Host string
		Port string
		Addr string
		ReadTimeout time.Duration
		WriteTimeout time.Duration
		IdleTimeout time.Duration
	}

	Database struct {
		Host string
		Port string
		User string
		Password string
		DBName string
		SSLMode string
		DSN string
	}

	JWTSecret string

	ProductClient struct {
		BaseURL string
	}

	Environment string
}

var Conf *Config

func Load() error {
	if err := godotenv.Load("../.env"); err != nil {
		return err
	}

	cfg := Config{}

	cfg.Server.Host = os.Getenv("HOST")
	cfg.Server.Port = os.Getenv("PORT")
	cfg.Server.ReadTimeout = time.Second * 15
	cfg.Server.WriteTimeout = time.Second * 30
	cfg.Server.IdleTimeout = time.Minute

	cfg.Server.Addr = fmt.Sprintf(`%s:%s`,
		cfg.Server.Host,
		cfg.Server.Port,
	)

	cfg.Database.Host = os.Getenv("DB_HOST")
	cfg.Database.Port = os.Getenv("DB_PORT")
	cfg.Database.User = os.Getenv("DB_USER")
	cfg.Database.Password = os.Getenv("DB_PASSWORD")
	cfg.Database.DBName = os.Getenv("DB_NAME")
	cfg.Database.SSLMode = os.Getenv("DB_SSLMODE")

	cfg.Database.DSN = fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=%s`,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	cfg.JWTSecret = os.Getenv("JWT_SECRET")

	cfg.ProductClient.BaseURL = os.Getenv("PRODUCT_URL")

	cfg.Environment = os.Getenv("ENV")

	Conf = &cfg

	return nil
}