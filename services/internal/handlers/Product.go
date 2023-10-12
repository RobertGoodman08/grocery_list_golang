package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"venv/services/db"
	"venv/services/scripts"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := db.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	scripts.RenderTemplate(w, "C:/Golang web/Grocery_Basket_Planner/services/internal/templates/index.html", products)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // Устанавливаем максимальный размер файла (10MB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error parsing form:", err)
		return
	}

	name := r.FormValue("name")

	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error converting quantity:", err)
		return
	}

	unitPrice, err := strconv.ParseFloat(r.FormValue("unit_price"), 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error converting unitPrice:", err)
		return
	}

	file, handler, err := r.FormFile("image_url")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error retrieving file:", err)
		return
	}
	defer file.Close()

	fileName := handler.Filename

	destFile, err := os.Create("C:\\Golang web\\Grocery_Basket_Planner\\services\\internal\\static\\images\\" + fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error creating destination file:", err)
		return
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error copying file:", err)
		return
	}

	quantityStr := strconv.Itoa(quantity)

	isChecked := r.FormValue("is_checked") == "on"

	err = db.AddProduct(name, "/static/images/"+fileName, quantityStr, unitPrice, isChecked)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error adding product:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/remove-from-cart/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	err = db.DeleteProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CheckedProducts(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/check_product/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	isChecked, err := strconv.ParseBool(r.FormValue("is_checked"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.CheckedProduct(id, isChecked)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
