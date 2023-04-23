package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Define the connection parameters for the database
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	// Connect to database
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	log.Println("Connected to database")

	// AutoMigrate the database
	err = db.AutoMigrate(&Product{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("Database migrated")

	// Create a new router
	r := mux.NewRouter()

	// Register endpoints
	r.HandleFunc("/products", createProduct(db)).Methods("POST")
	r.HandleFunc("/products/{id}", getProduct(db)).Methods("GET")
	r.HandleFunc("/products/{id}", updateProduct(db)).Methods("PUT")
	r.HandleFunc("/products/{id}", deleteProduct(db)).Methods("DELETE")
	r.HandleFunc("/products", getProducts(db)).Methods("GET")

	// Start the server
	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(":8081", r))
}
