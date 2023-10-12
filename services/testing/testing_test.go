package testing

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	"reflect"
	"testing"
	"venv/services/db"
	"venv/services/internal/models"
)

func TestGetProducts(t *testing.T) {
	dbe, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbe.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "image_url", "quantity", "unit_price"}).
		AddRow(1, "Product 1", "", 10, 5.0).
		AddRow(2, "Product 2", "", 8, 8.0)
	mock.ExpectQuery("SELECT id, name, image_url, quantity, unit_price FROM products").WillReturnRows(rows)

	dbConn, err := sql.Open("postgres", "")
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbConn.Close()

	dbe = dbConn

	products, err := db.GetProducts()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := []models.Product{
		{ID: 1, Name: "Product 1", ImageURL: "", Quantity: "2", UnitPrice: 5.0},
		{ID: 2, Name: "Product 2", ImageURL: "", Quantity: "8", UnitPrice: 8.0},
	}

	if !reflect.DeepEqual(products, expected) {
		t.Errorf("Expected %v, but got %v", expected, products)
	}
}
