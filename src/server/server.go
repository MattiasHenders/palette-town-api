package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/MattiasHenders/palette-town-api/config"
	h "github.com/MattiasHenders/palette-town-api/src/handlers"
	"github.com/MattiasHenders/palette-town-api/src/internal/middleware"
	s "github.com/MattiasHenders/palette-town-api/src/internal/server_helpers"
)

func Start(config *config.Config) {

	r := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{

		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},

		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// A good base middleware stack
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(chiMiddleware.Timeout(60 * time.Second))

	// Routes that require no authentication here...
	r.Group(func(r chi.Router) {

		// User authentication routes here...
		r.Post("/api/login", s.Handler(h.PostUserLoginHandler()))
		r.Post("/api/signup", s.Handler(h.PostUserSignupHandler()))
	})

	// Routes that require user authentication here...
	r.Group(func(r chi.Router) {

		// r.Use(middleware.Verifier(publicKeySet)) //TODO
		r.Use(middleware.Authenticator)
		r.Use(middleware.HydrateAuthUser())

		// User authentication routes here...
		r.Post("/api/logout", s.Handler(h.PostUserLogoutHandler()))

		// Colour routes here...
		r.Get("/api/colour/random", s.Handler(h.GetRandomColourPaletteHandler()))
		r.Get("/api/colour/colours", s.Handler(h.GetColourPromptColourPaletteHandler()))
		r.Get("/api/colour/words", s.Handler(h.GetWordPromptColourPaletteHandler()))
	})

	// Routes that require Admin access here...
	r.Group(func(r chi.Router) {
		// TODO: Authenticate the admin
		// r.Use(middleware.VerifyAdmin)

	})

	// Health checks here...
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	fmt.Printf("Running server on port %s...\n", config.Server.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", config.Server.Port), r)
}
