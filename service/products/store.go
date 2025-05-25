package products

import (
	"database/sql"

	"github.com/xudong7/ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}

	return products, nil
}

func scanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	p := new(types.Product)

	err := rows.Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Image,
		&p.Price,
		&p.Quantity,
		&p.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *Store) GetProductByName(name string) (*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE name = ?", name)
	if err != nil {
		return nil, err
	}

	p := new(types.Product)
	for rows.Next() {
		p, err = scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}
	}

	if p.ID == 0 {
		return nil, sql.ErrNoRows
	}

	return p, nil
}

func (s *Store) CreateProduct(p types.Product) error {
	_, err := s.db.Exec(
		"INSERT INTO products (name, description, image, price, quantity) VALUES (?, ?, ?, ?, ?)",
		p.Name,
		p.Description,
		p.Image,
		p.Price,
		p.Quantity,
	)
	if err != nil {
		return err
	}
	return nil
}
