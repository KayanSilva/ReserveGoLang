package models

import (
	"github.com/KayanSilva/ReserveGoLang/web-api/db"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetProducts() []Product {
	db := db.ConnectDB()
	defer db.Close()
	rows, err := db.Query("SELECT id, name, description, price, quantity FROM products ORDER BY name ASC")
	if err != nil {
		return nil
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity); err != nil {
			return nil
		}
		products = append(products, p)
	}

	return products
}

func CreateProduct(p Product) error {
	db := db.ConnectDB()
	defer db.Close()
	insertProduct, err := db.Prepare("INSERT INTO products(name, description, price, quantity) VALUES ($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	_, err = insertProduct.Exec(p.Name, p.Description, p.Price, p.Quantity)
	return err
}

func DeleteProduct(id int) error {
	db := db.ConnectDB()
	defer db.Close()
	deleteProduct, err := db.Prepare("DELETE FROM products WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	_, err = deleteProduct.Exec(id)
	return err
}

func GetProductByID(id int) (Product, error) {
	db := db.ConnectDB()
	defer db.Close()
	var p Product
	err := db.QueryRow("SELECT id, name, description, price, quantity FROM products WHERE id=$1", id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Quantity)
	if err != nil {
		return p, err
	}
	return p, nil
}

func UpdateProduct(p Product) error {
	db := db.ConnectDB()
	defer db.Close()
	updateProduct, err := db.Prepare("UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	_, err = updateProduct.Exec(p.Name, p.Description, p.Price, p.Quantity, p.ID)
	return err
}
