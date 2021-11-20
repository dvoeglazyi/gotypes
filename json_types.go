package gotypes

import (
	"encoding/base64"
	"errors"
)

// Base64 bytes slice that encodes / decodes in JSON in Base64.
// Срез байт, который кодируется или декодируется из JSON в формате Base64.
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
