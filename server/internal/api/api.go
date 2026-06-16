// Package api wires up the HTTP routes for the server.
//
// This is the base template with NO database. When you add a database
// template later, inject a Store implementation here and pass it to the
// handlers — do not put database logic directly in this file.
package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Router builds the application's HTTP handler with middleware and routes.
func Router() http.Handler {
	r := chi.NewRouter()

	// Standard middleware: request IDs, logging, panic recovery, timeouts.
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(15 * time.Second))

	// CORS so the Expo mobile app (running on a different origin during dev)
	// can call this API. Tighten AllowedOrigins for production.
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Liveness/health check.
	r.Get("/healthz", handleHealth)

	// Versioned API routes live under /api.
	r.Route("/api", func(r chi.Router) {
		r.Get("/hello", handleHello)
	})

	return r
}

// handleHealth reports that the server is up. Useful for load balancers,
// container orchestrators, and the mobile app's connectivity check.
func handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"status": "ok",
		"time":   time.Now().UTC().Format(time.RFC3339),
	})
}

// handleHello is a sample endpoint the mobile app calls to prove the
// end-to-end wiring works.
func handleHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"message": "Hello, " + name + "!",
	})
}

// writeJSON is a small helper to encode a value as a JSON response.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
