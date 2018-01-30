// handler package is defined handler.
// handler role is call form and hub.
// [enable dependency]: form, hub, message
package handler

import "github.com/go-chi/chi"

// Registe handler to mux
func Registe(r *chi.Mux) {
	r.Get("/task", taskHandler)
}
