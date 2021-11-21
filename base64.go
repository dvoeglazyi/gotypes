package gotypes

import (
	"database/sql"
	"database/sql/driver"
	"encoding/base64"
	"errors"
)

// Base64 bytes slice that uses Base64 format to be:
// * encoded / decoded in JSON
// * added to DB / received from DB
// Срез байт, который использует формат Base64:
// * при кодировании / декодировании в JSON
// * добавлении значения в БД / получении значения из БД
type Base64 []byte

func (b Base64) MarshalJSON() ([]byte, error) {
	if len(b) == 0 {
		return nil, nil
	}
	result := make([]byte, base64.StdEncoding.EncodedLen(len(b))+2)
	base64.StdEncoding.Encode(result[1:len(result)-1], b)
	result[0] = '"'
	result[len(result)-1] = '"'
	return result, nil
}

func (b *Base64) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		*b = nil
		return nil
	} else if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("invalid base64 string")
	}
	data = data[1 : len(data)-1]
	if len(data) == 0 {
		*b = nil
		return nil
	}
	result := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	length, err := base64.StdEncoding.Decode(result, data)
	if err != nil {
		return err
	}
	*b = result[:length]
	return nil
}

func (b Base64) Value() (driver.Value, error) {
	if len(b) == 0 {
		return nil, nil
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func (b *Base64) Scan(src interface{}) error {
	if src == nil {
		*b = nil
		return nil
	}
	var value sql.NullString
	if err := value.Scan(src); err != nil {
		return err
	} else if !value.Valid {
		*b = nil
		return nil
	}
	result, err := base64.StdEncoding.DecodeString(value.String)
	if err != nil {
		return err
	}
	*b = result
	return nil
}
