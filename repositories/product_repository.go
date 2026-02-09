package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
)

var ErrProductNotFound = errors.New("product not found")

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAll(name string) ([]models.ProductWithCategory, error) {
	query := `
		SELECT
			p.id,
			p.name,
			p.price,
			p.stock,
			p.category_id,
			c.name AS category_name
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
	`

	args := []interface{}{}
	if name != "" {
		query += " WHERE p.name ILIKE $1"
		args = append(args, "%"+name+"%")
	}

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []models.ProductWithCategory{}

	for rows.Next() {
		var p models.ProductWithCategory
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Price,
			&p.Stock,
			&p.CategoryID,
			&p.CategoryName,
		); err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (repo *ProductRepository) Create(p *models.Product) error {
	query := `
		INSERT INTO products (name, price, stock, category_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	return repo.db.QueryRow(
		query,
		p.Name,
		p.Price,
		p.Stock,
		p.CategoryID,
	).Scan(&p.ID)
}

func (repo *ProductRepository) GetById(id int) (*models.ProductWithCategory, error) {
	query := `
		SELECT
			p.id,
			p.name,
			p.price,
			p.stock,
			p.category_id,
			c.name AS category_name
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id
		WHERE p.id = $1
	`

	var p models.ProductWithCategory
	err := repo.db.QueryRow(query, id).Scan(
		&p.ID,
		&p.Name,
		&p.Price,
		&p.Stock,
		&p.CategoryID,
		&p.CategoryName,
	)

	if err == sql.ErrNoRows {
		return nil, ErrProductNotFound
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (repo *ProductRepository) Update(p *models.Product) error {
	query := `
		UPDATE products
		SET name = $1, price = $2, stock = $3, category_id = $4
		WHERE id = $5
	`

	result, err := repo.db.Exec(
		query,
		p.Name,
		p.Price,
		p.Stock,
		p.CategoryID,
		p.ID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrProductNotFound
	}

	return nil
}

func (repo *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"

	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrProductNotFound
	}

	return nil
}
