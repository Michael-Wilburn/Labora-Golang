package main

import (
	"labora/clases/Practicas/practica14_9-01-2023/movie-api/controllers"
	middlewares "labora/clases/Practicas/practica14_9-01-2023/movie-api/middlewares"
	services "labora/clases/Practicas/practica14_9-01-2023/movie-api/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Establecer conexión a la base de datos
	db, err := services.EstablishDbConnection()
	if err != nil {
		log.Fatal("Error establishing database connection:", err)
	}
	defer func() {
		if cerr := db.Close(); cerr != nil {
			log.Println("Error closing database connection:", cerr)
		}
	}()

	// Configurar controladores
	movieController := controllers.NewMovieController(db) // Asegúrate de tener un constructor adecuado

	// Configurar router con gorilla/mux
	router := mux.NewRouter()

	// Definir rutas utilizando el router de mux
	router.HandleFunc("/api/peliculas", movieController.AddMovie).Methods("POST")
	router.HandleFunc("/api/peliculas/{id}", movieController.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/peliculas/{id}", movieController.UpdateMovie).Methods("PUT")
	router.HandleFunc("/api/peliculas", movieController.GetMovies).Methods("GET")
	router.HandleFunc("/api/peliculas/{id}", movieController.GetMovieDetails).Methods("GET")
	router.HandleFunc("/api/peliculas/", movieController.GetMoviesByName).Methods("GET")

	//manejar los permisos CORS para todas las rutas
	corsHandler := middlewares.GetCorsHandler().Handler(router)

	// Iniciar el servidor con el router de mux
	log.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", corsHandler); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
