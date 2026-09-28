package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"catatan-backend/handler"
	"catatan-backend/repository"
	"catatan-backend/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Gagal membuka koneksi", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Gagal konek ke database:", err)
	}
	fmt.Println("Berhasil konek ke PostgreSQL!")

	noteRepo := repository.NewNoteRepository(db)
	noteService := service.NewNoteService(noteRepo)
	noteHandler := handler.NewNoteHandler(noteService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	http.HandleFunc("POST /categories/create", categoryHandler.Create)

	http.HandleFunc("POST /notes/create", noteHandler.Create)

	fmt.Println("Server jalan di http://localhost:8088")
	http.ListenAndServe(":8088", nil)
}
