package cmd

import (
	"hexagonal02/domain/usecases"
	api "hexagonal02/infraestructure/controllers"
	"hexagonal02/infraestructure/controllers/middleware"
	database "hexagonal02/infraestructure/database"
	"log"
	"net/http"
	"time"
)

func Start() {
	// Conectar a la base de datos
    //db,err := database.NewMySQLDB()
	db,err := database.NewSQLite3()

	if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()
	
	 // Inicializar el repositorio de usuario
	userRepository := database.NewUserRepositoryImpl(db)

	 // Inicializar el caso de uso de usuario
	userService := usecases.NewUserService(userRepository)

	userHandler := &api.UserHandler{UserService: *userService}

	// Inicializar el enrutador HTTP	
	router := api.NewRouter(userHandler)
	
	apiKeyMiddleware := middleware.NewApiKeyAuthMiddleware(router)

	srv := &http.Server{
        Handler:      apiKeyMiddleware,
        Addr:         "127.0.0.1:8000",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

	// Iniciar el servidor HTTP	

	log.Println("Starting server on :8000")	
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
