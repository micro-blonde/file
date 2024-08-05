package file

import "time"

type StoreRequest struct {
	Data      []byte
	Name      string
	Type      string
	BaseDir   *string
	ExpiresAt *time.Time
	AccountId *uint64
}

type StoreResponse[T Model] struct {
	*File[T]
	AbsPath string
	Url     string
}
