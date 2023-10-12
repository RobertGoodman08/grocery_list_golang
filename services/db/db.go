package db

import (
	"database/sql"
	"fmt"
	"log"
	"venv/services/internal/models"
)

var db *sql.DB

const (
	dbHost     = "localhost"
	dbPort     = your  port
	dbUser     = "your ursername"
	dbPassword = "your password"
	dbName     = "your dbname"
)

func init() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
}

func GetProducts() ([]models.Product, error) {
	rows, err := db.Query("SELECT id, name, image_url, quantity, unit_price, is_checked FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.ImageURL, &product.Quantity, &product.UnitPrice, &product.IsChecked)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func AddProduct(name string, imageURL string, quantity string, unitPrice float64, isChecked bool) error {
	_, err := db.Exec("INSERT INTO products (name, image_url, quantity, unit_price, is_checked) VALUES ($1, $2, $3, $4, $5)",
		name, imageURL, quantity, unitPrice, isChecked)
	return err
}

func DeleteProduct(id int) error {
	_, err := db.Exec("DELETE FROM products WHERE id = $1", id)
	return err
}

func CheckedProduct(id int, isChecked bool) error {
	_, err := db.Exec("UPDATE products SET is_checked = $1 WHERE id = $2", isChecked, id)
	return err
}
