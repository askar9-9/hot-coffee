package entity

type InventoryItem struct {
	IngredientID string  `json:"ingredient_id"`
	Name         string  `json:"name"`
	Quantity     float64 `json:"quantity"`
	Unit         string  `json:"unit"`
}

type InventoryTransaction struct {
	ID              string  `json:"transaction_id"`
	InventoryItemID string  `json:"inventory_item_id"`
	Quantity        float64 `json:"quantity"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}
