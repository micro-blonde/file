package file

import (
	"io"
	"time"
)

type StoreRequest struct {
	Reader    io.ReadSeeker
	Data      []byte
	Category  string
	Type      string
	Name      string
	Extension string
	BaseDir   *string
	ExpiresIn *time.Duration
	AccountId *uint64
}

type StoreResponse[T Model] struct {
	*File[T]
	AbsPath string
	Url     string
}
