package glusterfs_storage

import (
	"context"

	s3_storage "github.com/storage-lock/go-s3-storage"
	)

// GlusterfsStorage glusterfs对象存储实现 (S3兼容)
type GlusterfsStorage struct {
	*s3_storage.S3Storage
}

func NewGlusterfsStorage(ctx context.Context, options *s3_storage.S3StorageOptions) (*GlusterfsStorage, error) {
	s3Storage, err := s3_storage.NewS3Storage(ctx, options)
	if err != nil {
		return nil, err
	}
	return &GlusterfsStorage{S3Storage: s3Storage}, nil
}

const StorageName = "glusterfs-storage"

func (x *GlusterfsStorage) GetName() string {
	return StorageName
}

// Capabilities 声明存储支持的能力
