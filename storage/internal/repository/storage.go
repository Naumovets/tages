package repository

import (
	"log/slog"

	"github.com/Naumovets/tages/internal/entity"
)

func (r *Repository) Create(file *entity.File) error {
	_, err := r.db.Model(file).Insert()
	if err != nil {
		slog.Debug("failed to create a file", "error", err)

		return err
	}

	slog.Debug("file created", "file", file)

	return nil
}

func (r *Repository) GetList(limit, offset int) ([]entity.File, error) {
	var files []entity.File
	err := r.db.Model(&files).Limit(limit).Offset(offset).Select()
	if err != nil {
		slog.Debug("failed to list files", "error", err)

		return nil, err
	}

	return files, nil
}

func (r *Repository) GetById(id string) (*entity.File, error) {
	var file entity.File
	err := r.db.Model(&file).Where("id = ?", id).Select()
	if err != nil {
		slog.Debug("failed to get a file", "error", err)

		return nil, err
	}

	return &file, err
}
