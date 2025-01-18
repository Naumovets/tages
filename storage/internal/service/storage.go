package service

import (
	"fmt"
	"io"
	"log/slog"
	"time"

	"github.com/Naumovets/tages/internal/entity"
	tages "github.com/Naumovets/tages/pkg/proto/storage"
	"github.com/google/uuid"
)

// Upload handles the gRPC stream for uploading files. It creates a new file,
// writes chunks received from the client to it and saves the file to the
// repository. It returns the ID of the uploaded file or an error if something
// fails.
func (s *service) Upload(stream tages.Storage_UploadServer) (string, error) {
	file := entity.NewFile()
	defer func() {
		if err := file.OutputFile.Close(); err != nil {
			slog.Debug("failed to close file", "error", err)
		}
	}()

	file.CreatedAt = time.Now().Format(time.DateTime)
	file.UpdatedAt = time.Now().Format(time.DateTime)
	file.Id = uuid.New().String()

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			slog.Debug("failed to receive a chunk", "error", err)

			return "", err
		}

		file.FileName = req.GetFileName()
		fmt.Println(file.FileName)

		if file.FilePath == "" {
			file.SetFile(s.cfg.Storage.Location)
		}

		chunk := req.GetChunk()
		if err := file.Write(chunk); err != nil {
			slog.Debug("failed to write a chunk", "error", err)

			return "", err
		}
	}

	err := s.rep.Create(file)
	if err != nil {
		slog.Debug("failed to create a file", "error", err)
		fmt.Println(err)

		return "", err
	}

	return file.Id, nil
}

// GetList returns a list of files from the repository, starting from the given
// offset and limited to the given limit. It returns an error if something fails.
func (s *service) GetList(limit, offset int64) ([]*tages.File, error) {
	files, err := s.rep.GetList(int(limit), int(offset))
	if err != nil {
		slog.Debug("failed to list files", "error", err)

		return nil, err
	}

	var tagesFiles []*tages.File
	for _, file := range files {
		tagesFiles = append(tagesFiles, &tages.File{
			Id:        file.Id,
			FileName:  file.FileName,
			CreatedAt: file.CreatedAt,
			UpdatedAt: file.UpdatedAt,
		})
	}

	return tagesFiles, nil
}

// Download returns a file from the repository by its id. It returns an error if
// something fails.
func (s *service) Download(id string, stream tages.Storage_DownloadServer) error {
	fileDB, err := s.rep.GetById(id)
	if err != nil {
		slog.Debug("failed to get a file", "error", err)

		return err
	}

	file, err := fileDB.Get(s.cfg.Storage.Location)
	if err != nil {
		slog.Debug("failed to get a file", "error", err)

		return err
	}

	buf := make([]byte, s.cfg.Storage.BatchSize)
	batchNumber := 1
	for {
		num, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		chunk := buf[:num]

		if err := stream.Send(&tages.DownloadResponse{FileName: fileDB.FileName, Chunk: chunk}); err != nil {
			return err
		}
		batchNumber += 1

	}

	if err := file.Close(); err != nil {
		slog.Debug("failed to close file", "error", err)

		return err
	}

	return nil
}
