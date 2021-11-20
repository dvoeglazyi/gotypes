package gotypes

import (
	"database/sql"
	"database/sql/driver"
)

// NullUint unsigned integer. Value 0 corresponds to db-value NULL.
// Беззнаковое целое. Значение 0 соответствует значению NULL в БД.
type NullUint uint

func (nu NullUint) Value() (driver.Value, error) {
	if nu == 0 {
		return nil, nil
	}
	return int64(nu), nil
}

func (nu *NullUint) Scan(src interface{}) error {
	if src == nil {
		*nu = 0
		return nil
	}
	var value sql.NullInt64
	if err := value.Scan(src); err != nil {
		return err
	} else if !value.Valid {
		*nu = 0
		return nil
	}
	*nu = NullUint(value.Int64)
	return nil
}

// NullString string. Empty value '' corresponds to db-value NULL.
// Строка. Пустое значение '' соответствует значению NULL в БД.
type NullString string

func (ns NullString) Value() (driver.Value, error) {
	if ns == "" {
		return nil, nil
	}
	return string(ns), nil
}

func (ns *NullString) Scan(src interface{}) error {
	if src == nil {
		*ns = ""
		return nil
	}
	var value sql.NullString
	if err := value.Scan(src); err != nil {
		return err
	} else if !value.Valid {
		*ns = ""
		return nil
	}
	*ns = NullString(value.String)
	return nil
}
