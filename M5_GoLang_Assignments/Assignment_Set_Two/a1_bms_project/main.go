package main

import (
	db "blog/config"
	"blog/controller"
	"blog/middleware"
	"blog/repository"
	"blog/service"
	"database/sql"
	"fmt"
	"net/http"

	_ "modernc.org/sqlite" 
)

func InitializeAuthDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./myblogs.db")
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to the database successfully.")

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL,
	password TEXT NOT NULL
	);`
	if _, err := db.Exec(createTableSQL); err != nil {
		return nil, fmt.Errorf("failed to create table: %v", err)
	}

	insertSQL := `INSERT INTO users (username, password) VALUES (?,?);`
	if _, err := db.Exec(insertSQL, "admin", "admin"); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	auth_db, err := InitializeAuthDatabase()
	if err != nil {
		fmt.Println("Auth Database initialization failed:", err)
		return
	}
	defer auth_db.Close()

	db.InitializeDatabase()

	blogRepo := repository.NewBlogRepository(db.GetDB())
	blogService := service.NewBlogService(blogRepo)
	blogController := controller.NewBlogController(blogService)

	mux := http.NewServeMux()

	mux.HandleFunc("/blogs", blogController.GetAllBlogs)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			blogController.CreateBlog(w, r)
		default:
			http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/blog/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			blogController.GetBlog(w, r)
		case http.MethodPut:
			blogController.UpdateBlog(w, r)
		case http.MethodDelete:
			blogController.DeleteBlog(w, r)
		default:
			http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
		}
	})

	loggedMux := middleware.LoggingMiddleware(mux)
	authenticatedMux := middleware.Authmiddleware(loggedMux)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", authenticatedMux); err != nil {
		fmt.Println("Error Starting Server:", err)
	}
}
