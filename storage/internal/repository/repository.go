package repository

import (
	"github.com/Naumovets/tages/internal/entity"
	"github.com/go-pg/pg/v10"
)

type IRepository interface {
	Create(file *entity.File) error
	GetList(limit, offset int) ([]entity.File, error)
	GetById(id string) (*entity.File, error)
}

type Repository struct {
	db *pg.DB
}

func NewRepository(db *pg.DB) *Repository {
	return &Repository{
		db: db,
	}
}
