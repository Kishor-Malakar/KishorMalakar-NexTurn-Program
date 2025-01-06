package repository

import (
	"database/sql"
	"ecommerce/model"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (repo *ProductRepository) CreateProduct(product *model.Product) (*model.Product, error) {
	stmt, err := repo.DB.Prepare("INSERT INTO products (name, decription, price, stock, category_id) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(product.Name, product.Decription, product.Price, product.Stock, product.Category_ID)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	product.ID = int(id)
	return product, nil
}

func (repo *ProductRepository) GetProduct(id int) (*model.Product, error) {
	row := repo.DB.QueryRow("SELECT * FROM products WHERE id = ?", id)
	product := &model.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Decription, &product.Price, &product.Stock, &product.Category_ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (repo *ProductRepository) GetAllProducts(page, limit int) ([]model.Product, error) {
	offset := (page - 1) * limit
	query := "SELECT * FROM products LIMIT ? OFFSET ?"

	rows, err := repo.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Decription, &product.Price, &product.Stock, &product.Category_ID)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (repo *ProductRepository) UpdateProduct(product *model.Product) (*model.Product, error) {
	stmt, err := repo.DB.Prepare("UPDATE products SET stock = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Stock, product.ID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (repo *ProductRepository) DeleteProduct(id int) error {
	stmt, err := repo.DB.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
