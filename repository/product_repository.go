package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return ProductRepository{
		connection: db,
	}
}

func (r *ProductRepository) GetAllProducts() ([]model.GetProduct, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := r.connection.Query(query) 
	if err != nil {
		fmt.Println("Error executing query:", err)
		return []model.GetProduct{}, err
	}

	var products []model.GetProduct
	var product model.GetProduct

	for rows.Next(){
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return []model.GetProduct{}, err
		}
		products = append(products, product)
	}
	rows.Close()
	return products, nil
}

func (r *ProductRepository) PostCreateProduct(product model.PostProduct) (model.PostProduct, error) {
	fmt.Println("Product Name:", product.Name, product.Price)
	query := "INSERT INTO product (product_name, price) VALUES ($1, $2)"
	_, err := r.connection.Exec(query, product.Name, product.Price)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return model.PostProduct{}, err
	}
	return product, nil
}

func (r *ProductRepository) DeleteProduct(id int) error {
	query := "DELETE FROM product WHERE id = $1"
	_, err := r.connection.Exec(query, id)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return err
	}
	return nil
}