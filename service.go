package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func createProduct(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product Product
		json.NewDecoder(r.Body).Decode(&product)

		db.Create(&product)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(product)
	}
}

func getProducts(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var products []Product
		db.Find(&products)

		json.NewEncoder(w).Encode(products)
	}
}

func getProduct(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		productID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		var product Product
		result := db.First(&product, productID)
		if result.Error != nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(product)
	}
}

func updateProduct(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		productID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		var product Product
		result := db.First(&product, productID)
		if result.Error != nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		json.NewDecoder(r.Body).Decode(&product)

		db.Save(&product)

		json.NewEncoder(w).Encode(product)
	}
}

func deleteProduct(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		productID, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		var product Product
		result := db.First(&product, productID)
		if result.Error != nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		db.Delete(&product)

		w.WriteHeader(http.StatusOK)
	}
}
