package storage

import (
	"bytes"
)

type ImageStorage interface {
	Download(key string) (*bytes.Buffer, error)
}
