package controller

import (
	"context"

	tages "github.com/Naumovets/tages/pkg/proto/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Upload handles the gRPC stream for uploading files. It limits concurrent uploads
func (s *serverStorage) Upload(stream tages.Storage_UploadServer) error {
	s.uploadDownloadLimiter <- struct{}{}
	defer func() { <-s.uploadDownloadLimiter }()

	id, err := s.service.Upload(stream)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return stream.SendAndClose(&tages.UploadResponse{Id: id})
}

// GetList handles the gRPC call for listing files. It limits concurrent list requests
func (s *serverStorage) GetList(ctx context.Context, req *tages.ListFilesRequest) (*tages.ListFilesResponse, error) {
	s.getListLimiter <- struct{}{}
	defer func() { <-s.getListLimiter }()

	if req.GetLimit() == 0 {
		req.Limit = 10
	}

	if req.GetOffset() == 0 {
		req.Offset = 0
	}

	files, err := s.service.GetList(req.GetLimit(), req.GetOffset())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &tages.ListFilesResponse{Files: files}, nil
}

// Download handles the gRPC stream for downloading files. It limits concurrent downloads
func (s *serverStorage) Download(req *tages.DownloadRequest, stream tages.Storage_DownloadServer) error {
	s.uploadDownloadLimiter <- struct{}{}
	defer func() { <-s.uploadDownloadLimiter }()

	err := s.service.Download(req.GetId(), stream)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
