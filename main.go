package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"catatan-backend/handler"
	"catatan-backend/repository"
	"catatan-backend/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

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
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	fmt.Println("Berhasil konek ke PostgreSQL!")

	noteRepo := repository.NewNoteRepository(db)
	noteService := service.NewNoteService(noteRepo)
	noteHandler := handler.NewNoteHandler(noteService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	http.HandleFunc("POST /categories", categoryHandler.Create)
	http.HandleFunc("GET /categories", categoryHandler.GetAll)
	http.HandleFunc("GET /categories/{id}", categoryHandler.GetByID)
	http.HandleFunc("DELETE /categories/{id}", categoryHandler.Delete)
	http.HandleFunc("PUT /categories/{id}", categoryHandler.Update)

	http.HandleFunc("POST /notes", noteHandler.Create)
	http.HandleFunc("GET /notes", noteHandler.GetAll)
	http.HandleFunc("GET /notes/{id}", noteHandler.GetByID)
	http.HandleFunc("DELETE /notes/{id}", noteHandler.Delete)
	http.HandleFunc("PUT /notes/{id}", noteHandler.Update)

	fmt.Println("Server jalan di http://localhost:8088")
	log.Fatal(http.ListenAndServe(":8088", enableCORS(http.DefaultServeMux)))
}
