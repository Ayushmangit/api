package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Ayushmangit/api/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// making a pass around config so the handlers get the access top the database
type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the env")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the env ")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("connection to the database failed")
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"http://*",
			"https://*",
		},
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{
			"Link",
		},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	v1Router := chi.NewRouter()
	v1Router.Get("/health", HandlerReadiness)
	v1Router.Get("/err", HandlerError)
	v1Router.Get("/users", apiCfg.HandlerGetAllUsers)
	v1Router.Post("/users", apiCfg.HandlerCreateUser)
	v1Router.Delete("/users/{id}", apiCfg.HandlerDestroyUser)
	v1Router.Put("/users/{id}", apiCfg.HandlerUpdateUser)

	router.Mount("/v1", v1Router)

	log.Printf("Server starting on port %v", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
