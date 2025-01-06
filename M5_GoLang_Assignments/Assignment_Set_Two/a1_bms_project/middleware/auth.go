package middleware

import (
	"database/sql"
	"encoding/base64"
	"net/http"
	"strings"
)

func Authmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(authHeader, "Basic "))
		if err != nil {
			http.Error(w, "Invalid Authorization Header", http.StatusUnauthorized)
			return
		}
		credentials := strings.SplitN(string(payload), ":", 2)
		if len(credentials) != 2 {
			http.Error(w, "Invalid credentials format", http.StatusUnauthorized)
			return
		}

		// Validate credentials against the database
		username := credentials[0]
		password := credentials[1]
		db, err := sql.Open("sqlite", "myblogs.db")
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		var dbPassword string
		query := "SELECT password FROM users WHERE username = ?"
		err = db.QueryRow(query, username).Scan(&dbPassword)
		if err != nil || dbPassword != password {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
