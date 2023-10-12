package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"venv/services/internal/handlers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Определение маршрутов
	router.HandleFunc("/", handlers.GetProducts).Methods("GET")
	router.HandleFunc("/add-product", handlers.AddProduct).Methods("POST")
	router.HandleFunc("/remove-from-cart/{id}", handlers.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/check_product/{id}", handlers.CheckedProducts).Methods("PUT")

	//  статические файлы
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("internal/static"))))

	return router
}
