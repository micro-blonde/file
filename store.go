package file

import "time"

type StoreRequest struct {
	Data      []byte
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
