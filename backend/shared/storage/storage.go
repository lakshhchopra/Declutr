package storage

import (
	"context"
	"io"
)

type StorageProvider interface {
	PutObject(ctx context.Context, key string, data io.Reader, size int64, contentType string) error
	GetObject(ctx context.Context, key string) (io.ReadCloser, error)
	DeleteObject(ctx context.Context, key string) error
	GeneratePreSignedURL(ctx context.Context, key string, expirationSeconds int) (string, error)
}

type LocalStorageProvider struct {
	BaseDir string
}

func NewLocalStorageProvider(baseDir string) *LocalStorageProvider {
	return &LocalStorageProvider{BaseDir: baseDir}
}

func (l *LocalStorageProvider) PutObject(ctx context.Context, key string, data io.Reader, size int64, contentType string) error {
	return nil
}

func (l *LocalStorageProvider) GetObject(ctx context.Context, key string) (io.ReadCloser, error) {
	return nil, nil
}

func (l *LocalStorageProvider) DeleteObject(ctx context.Context, key string) error {
	return nil
}

func (l *LocalStorageProvider) GeneratePreSignedURL(ctx context.Context, key string, expirationSeconds int) (string, error) {
	return "http://localhost:8080/storage/mock/" + key, nil
}
