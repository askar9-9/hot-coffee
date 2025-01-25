package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type MenuItem struct {
	ID          string      `json:"product_id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Size        string      `json:"size"`
	Category    string      `json:"category"`
	Tags        []string    `json:"tags"`
	MetaData    MetaDataMap `json:"meta_data"`
}

type MenuItemIngredient struct {
	ID           string  `json:"id"`
	MenuItemID   string  `json:"product_id"`
	IngredientID string  `json:"ingredient_id"`
	Quantity     float64 `json:"quantity"`
}

type MetaDataMap map[string]interface{}

func (p MetaDataMap) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

func (p *MetaDataMap) Scan(src interface{}) error {
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