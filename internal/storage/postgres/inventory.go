package postgres

import "hot-coffee/internal/entity"

func (p *Postgres) GetAllInventoryItems() ([]*entity.InventoryItem, error) {
	sql := `SELECT * FROM inventory`

	rows, err := p.db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var InventoryItems []*entity.InventoryItem
	for rows.Next() {
		var InventoryItem entity.InventoryItem
		err := rows.Scan(
			&InventoryItem.IngredientID,
			&InventoryItem.Name,
			&InventoryItem.Quantity,
			&InventoryItem.Unit,
			&InventoryItem.Price,
		)

		if err != nil {
			return nil, err
		}
		InventoryItems = append(InventoryItems, &InventoryItem)
	}

	return InventoryItems, nil
}

func (p *Postgres) GetInventoryItemByID(id string) (*entity.InventoryItem, error) {
	sql := `SELECT id, name, quantity, unit, price FROM inventory WHERE id = $1`

	var InventoryItem entity.InventoryItem
	err := p.db.QueryRow(sql, id).Scan(
		&InventoryItem.IngredientID,
		&InventoryItem.Name,
		&InventoryItem.Quantity,
		&InventoryItem.Unit,
		&InventoryItem.Price,
	)

	if err != nil {
		return nil, err
	}

	return &InventoryItem, nil
}

func (p *Postgres) AddInventoryItem(item *entity.InventoryItem) error {
	sql := `INSERT INTO inventory (id, name, quantity, unit, price) VALUES ($1, $2, $3, $4, $5)`

	_, err := p.db.Exec(sql, item.IngredientID, item.Name, item.Quantity, item.Unit, item.Price)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) UpdateInventoryItemByID(id string, item *entity.InventoryItem) error {
	sql := `UPDATE inventory SET name = $1, quantity = $2, unit = $3, price = $4 WHERE id = $5`

	_, err := p.db.Exec(sql, item.Name, item.Quantity, item.Unit, item.Price, item.IngredientID)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) DeleteInventoryItemByID(id string) error {
	sql := `DELETE FROM inventory WHERE id = $1`

	_, err := p.db.Exec(sql, id)
	if err != nil {
		return err
	}

	return nil
}
