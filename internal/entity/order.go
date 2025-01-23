package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Order struct {
	ID          string  `json:"order_id"`
	CustomerID  string  `json:"customer_id"`
	Status      string  `json:"status"`
	TotalAmount float64 `json:"total_amount"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type OrderItems struct {
	ID            string    `json:"order_item_id"`
	OrderID       string    `json:"order_id"`
	ProductID     string    `json:"product_id"`
	Quantity      int       `json:"quantity"`
	Price         int       `json:"price"`
	Customization CustomMap `json:"customization"`
}

type OrderStatusHistory struct {
	ID          string `json:"status_id"`
	OrderID     string `json:"order_id"`
	Status      string `json:"status"`
	ChangedDate string `json:"changed_date"`
}

type CustomMap map[string]interface{}

func (c CustomMap) Value() (driver.Value, error) {
	j, err := json.Marshal(c)
	return j, err
}

func (c *CustomMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*c, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("Type assertion .(map[string]interface{}) failed.")
	}

	return nil
}
