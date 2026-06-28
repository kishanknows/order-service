package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kishanknows/order-service/internal/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	db, err := sql.Open("postgres", config.Conf.Database.DSN)

	if err != nil {
		return err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)

	if err := db.Ping(); err != nil {
		return err
	}

	fmt.Println("DB connected successfully")
	DB = db
	
	return nil
}