package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/rodrigueslg/codedu-goexpert/rest-api/configs"
	"github.com/rodrigueslg/codedu-goexpert/rest-api/internal/entity"
	"github.com/rodrigueslg/codedu-goexpert/rest-api/internal/handlers"
	"github.com/rodrigueslg/codedu-goexpert/rest-api/internal/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	mymiddleware "github.com/rodrigueslg/codedu-goexpert/rest-api/internal/middleware"
	swagger "github.com/swaggo/http-swagger"

	_ "github.com/rodrigueslg/codedu-goexpert/rest-api/api"
)

// @title CodeEdu GoExpert - REST API
// @version 1.0
// @description Product API with bearer token authentication
// @termsOfService http://swagger.io/terms/

// @contact.name Luis Rodrigues
// @contact.url http://github.com/rodrigueslg
// @contact.email rodrigueslg@outlook.com

// @license.name MIT
// @license.url http://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&entity.Product{}, &entity.User{})
	if err != nil {
		panic(err)
	}

	productRepo := repository.NewProductRepository(db)
	userRepo := repository.NewUserRepository(db)

	productHandler := handlers.NewProductHandler(productRepo)
	userHandler := handlers.NewUserHandler(userRepo, configs.TokenAuth, configs.JWTExpiresIn)

	r := chi.NewRouter()
	//r.Use(middleware.Logger)
	r.Use(mymiddleware.LogRequest)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetProducts)
		r.Post("/", productHandler.CreateProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Post("/auth", userHandler.Auth)
	})

	r.Get("/docs/*", swagger.Handler(swagger.URL("http://localhost:8080/docs/doc.json")))

	log.Fatal(http.ListenAndServe(":8080", r))
}
