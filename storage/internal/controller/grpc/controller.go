package controller

import (
	"github.com/Naumovets/tages/config"
	"github.com/Naumovets/tages/internal/repository"
	"github.com/Naumovets/tages/internal/service"
	tages "github.com/Naumovets/tages/pkg/proto/storage"
)

type serverStorage struct {
	tages.UnimplementedStorageServer
	service               service.IService
	uploadDownloadLimiter chan struct{}
	getListLimiter        chan struct{}
}

func NewServerStorage(cfg *config.Config, rep repository.IRepository) *serverStorage {
	return &serverStorage{
		service:               service.NewService(cfg, rep),
		uploadDownloadLimiter: make(chan struct{}, 10),
		getListLimiter:        make(chan struct{}, 100),
	}
}
