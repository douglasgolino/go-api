package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProductById(id_product int) (*model.Product, error) {

	query, err := pr.connection.Prepare("SELECT id, product_name, price FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var product model.Product

	err = query.QueryRow(id_product).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &product, nil

}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {

	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.Id,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {

	var id int
	query, err := pr.connection.Prepare("INSERT INTO product (product_name, price) VALUES ($1,$2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (pr *ProductRepository) UpdateProduct(productId int, product model.Product) (model.Product, error) {

	query := "UPDATE product SET product_name = $1, price = $2 WHERE id = $3 RETURNING id"

	var updatedId int
	err := pr.connection.QueryRow(query, product.Name, product.Price, productId).Scan(&updatedId)
	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}

	// Verifica se o ID retornado corresponde ao ID que foi atualizado
	if updatedId != productId {
		return model.Product{}, fmt.Errorf("incompatibilidade de ID do produto: esperado %d mas retornou %d", productId, updatedId)
	}

	// Recupera o produto atualizado do banco de dados
	query = "SELECT id, product_name, price FROM product WHERE id = $1"
	var updatedProduct model.Product
	err = pr.connection.QueryRow(query, productId).Scan(&updatedProduct.Id, &updatedProduct.Name, &updatedProduct.Price)
	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}

	return updatedProduct, nil
}

func (pr *ProductRepository) DeleteProduct(id_product int) (int, error) {

	query := "DELETE FROM product WHERE id = $1"

	result, err := pr.connection.Exec(query, id_product)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return int(rowsAffected), nil
}
