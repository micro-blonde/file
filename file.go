package file

import (
	"time"
)

type File[T Model] struct {
	Id             string
	Key            string
	CreatedAt      time.Time
	ExpirationTime *time.Time
	// AccountId issuer of file
	AccountId *uint64
	// Type is MIME type according to file extension
	Type string
	// Name of file
	Name string
	// Extension of file
	Extension string
	// Path as relative path of file
	Path string
	// Meta is metadata of file containing more information about file
	Meta *Meta

	T T `gorm:"embedded" json:",inline"`
}
