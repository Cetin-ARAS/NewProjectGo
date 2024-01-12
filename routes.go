// routes.go

package routes

import (
	"github.com/go-chi/chi"
	"github.com/yourusername/yourprojectname/main"
)

// RegisterUserRoutes registers user-related routes
func RegisterUserRoutes(r *chi.Mux) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", main.GetAllUsers)
		r.Get("/{id}", main.GetUser)
		r.Post("/", main.CreateUser)
		r.Put("/{id}", main.UpdateUser)
		r.Delete("/{id}", main.DeleteUser)
	})
}
