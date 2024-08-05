package client

import (
	"time"

	"github.com/micro-blonde/file"
	pf "github.com/micro-blonde/file/proto/file"
)

func getFile[T file.Model](f *pf.File) *file.File[T] {
	r := &file.File[T]{
		Id:        f.Id,
		CreatedAt: time.UnixMicro(f.CreatedAt),
		Type:      f.Type,
		Category:  f.Category,
		Key:       f.Key,
		Name:      f.Name,
		Path:      f.Path,
	}
	if f.ExpirationTime != nil {
		t := f.ExpirationTime.AsTime()
		r.ExpirationTime = &t
	}
	if f.AccountId != 0 {
		r.AccountId = &f.AccountId
	}
	return r
}
