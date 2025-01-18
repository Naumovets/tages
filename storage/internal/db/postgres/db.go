package postgres

import (
	"fmt"

	"github.com/Naumovets/tages/config"
	"github.com/go-pg/pg/v10"
)

func NewConn(cfg *config.Config) (*pg.DB, error) {

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.DB.DB_HOST, cfg.DB.DB_PORT),
		User:     cfg.DB.USER,
		Password: cfg.DB.PASSWORD,
		Database: cfg.DB.DB_NAME,
	})

	return db, nil
}
