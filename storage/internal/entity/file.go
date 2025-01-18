package entity

import (
	"bytes"
	"os"
	"path/filepath"
)

type File struct {
	Id         string
	FileName   string
	CreatedAt  string
	UpdatedAt  string
	FilePath   string        `pg:"-"`
	buffer     *bytes.Buffer `pg:"-"`
	OutputFile *os.File      `pg:"-"`
}

func NewFile() *File {
	return &File{
		buffer: &bytes.Buffer{},
	}
}

func (f *File) SetFile(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	f.FilePath = filepath.Join(path, f.CreatedAt+f.FileName)

	file, err := os.Create(f.FilePath)
	if err != nil {
		return err
	}

	f.OutputFile = file

	return nil
}

func (f *File) Write(chunk []byte) error {
	if f.OutputFile == nil {
		return nil
	}

	_, err := f.OutputFile.Write(chunk)

	return err
}

func (f *File) Get(path string) (*os.File, error) {
	file, err := os.Open(filepath.Join(path, f.CreatedAt+f.FileName))
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (f *File) Close() error {
	return f.OutputFile.Close()
}
