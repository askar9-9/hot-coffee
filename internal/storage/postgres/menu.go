package postgres

import "hot-coffee/internal/entity"

func (p *Postgres) GetAllMenuItems() ([]*entity.MenuItem, error) {
	sql := `SELECT * FROM menu_items`

	rows, err := p.db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var menuItems []*entity.MenuItem
	for rows.Next() {
		var menuItem entity.MenuItem
		err := rows.Scan(
			&menuItem.ID,
			&menuItem.Name,
			&menuItem.Description,
			&menuItem.Price,
			&menuItem.Size,
			&menuItem.Category,
			&menuItem.Tags,
			&menuItem.MetaData,
		)

		if err != nil {
			return nil, err
		}
		menuItems = append(menuItems, &menuItem)
	}

	return menuItems, nil
}

func (p *Postgres) GetMenuItemByID(id int) (*entity.MenuItem, error) {
	sql := `SELECT * FROM menu_items WHERE id = $1`

	var menuItem entity.MenuItem
	err := p.db.QueryRow(sql, id).Scan(
		&menuItem.ID,
		&menuItem.Name,
		&menuItem.Description,
		&menuItem.Price,
		&menuItem.Size,
		&menuItem.Category,
		&menuItem.Tags,
		&menuItem.MetaData,
	)

	if err != nil {
		return nil, err
	}

	return &menuItem, nil
}

func (p *Postgres) CreateMenuItem(item *entity.MenuItem) error {
	sql := `INSERT INTO menu_items (product_id, name, description, price, size, category, tags, meta_data) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := p.db.Exec(sql, item.ID, item.Name, item.Description, item.Price, item.Size, item.Category, item.Tags, item.MetaData)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) UpdateMenuItem(item *entity.MenuItem) error {
	sql := `UPDATE menu_items SET name = $1, description = $2, price = $3, size = $4, category = $5, tags = $6, meta_data = $7 WHERE product_id = $8`

	_, err := p.db.Exec(sql, item.Name, item.Description, item.Price, item.Size, item.Category, item.Tags, item.MetaData, item.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) DeleteMenuItem(id int) error {
	sql := `DELETE FROM menu_items WHERE id = $1`

	_, err := p.db.Exec(sql, id)
	if err != nil {
		return err
	}

	return nil
}
