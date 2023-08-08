package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type DB struct {
	Pool *sql.DB
}

func New() (*DB, error) {
	return newFromURL(viper.GetString("database_uri"))
}

func newFromURL(url string) (*DB, error) {
	fmt.Println(url)
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	d := &DB{Pool: db}
	return d, nil
}
