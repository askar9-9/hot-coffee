package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Customer struct {
	ID          string  `json:"customer_id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Preferences PrefMap `json:"preferences"`
}

type PrefMap map[string]interface{}

func (p PrefMap) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

func (p *PrefMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*p, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("Type assertion .(map[string]interface{}) failed.")
	}

	return nil
}
