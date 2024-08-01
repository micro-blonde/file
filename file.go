package file

import "time"

type File[T Model] struct {
	Id             string
	CreatedAt      time.Time
	ExpirationTime *time.Time
	// Category of file
	Category string
	// AccountId issuer of file
	AccountId *uint64
	// Name of file
	Name string
	// Type is MIME type according to file extension
	Type string
	// Extension of file
	Extension string
	// Path as relative path of file
	Path string

	T T `gorm:"embedded" json:",inline"`
}
