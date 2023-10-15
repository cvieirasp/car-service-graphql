package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db   *sql.DB
	ID   string
	Name string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{
		db: db,
	}
}

func (c *Category) Create(name string) (*Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name) VALUES ($1, $2)", id, name)
	if err != nil {
		return nil, err
	}
	return &Category{ID: id, Name: name}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := []Category{}
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (c *Category) FindByID(id string) (Category, error) {
	var category Category
	if err := c.db.QueryRow("SELECT id, name FROM categories WHERE id = $1", id).Scan(&category.ID, &category.Name); err != nil {
		return Category{}, err
	}
	return category, nil
}

func (c *Category) FindByServiceID(serviceID string) (Category, error) {
	var category Category
	if err := c.db.QueryRow("SELECT c.id, c.name FROM categories c JOIN services s ON c.id = s.category_id WHERE s.id = $1", serviceID).Scan(&category.ID, &category.Name); err != nil {
		return Category{}, err
	}
	return category, nil
}
