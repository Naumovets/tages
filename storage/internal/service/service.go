package service

import (
	"github.com/Naumovets/tages/config"
	"github.com/Naumovets/tages/internal/repository"
	tages "github.com/Naumovets/tages/pkg/proto/storage"
)

type IService interface {
	Upload(stream tages.Storage_UploadServer) (string, error)
	GetList(limit, offset int64) ([]*tages.File, error)
	Download(id string, stream tages.Storage_DownloadServer) error
}

type service struct {
	cfg *config.Config
	rep repository.IRepository
}

func NewService(cfg *config.Config, rep repository.IRepository) IService {
	return &service{
		cfg: cfg,
		rep: rep,
	}
}
