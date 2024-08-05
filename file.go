package file

import (
	"time"
)

type File[T Model] struct {
	Id             string
	Key            string
	CreatedAt      time.Time
	ExpirationTime *time.Time
	// Type is MIME type according to file extension
	Type string
	// Category of file to be able to group each kind
	// of files existing in system(e.g. for oauth profile -> profile_picture)
	Category string
	// AccountId issuer of file
	AccountId *uint64
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
