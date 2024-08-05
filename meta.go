package file

import (
	"database/sql/driver"
	"encoding/json"
)

type Meta struct {
}

func (t Meta) Value() (driver.Value, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (t *Meta) Scan(src interface{}) error {
	data, ok := src.([]byte)
	if !ok {
		return nil
	}
	d := new(Meta)
	err := json.Unmarshal(data, d)
	if err != nil {
		return err
	}
	*t = *d
	return nil
}
