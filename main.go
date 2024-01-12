// main.go

package main

import (
	"log"
	"net/http"

	"github.com/Cetin-ARAS/NewProjectGo/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

// User model
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	var err error
	db, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// routes.go dosyasındaki RegisterUserRoutes fonksiyonunu çağır
	routes.RegisterUserRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}

// ... (diğer fonksiyonlar)
