package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	models "labora/clases/Practicas/practica14_9-01-2023/movie-api/models"
	services "labora/clases/Practicas/practica14_9-01-2023/movie-api/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MovieController struct {
	DB *sql.DB
}

func NewMovieController(db *sql.DB) *MovieController {
	return &MovieController{DB: db}
}

func (mc *MovieController) AddMovie(w http.ResponseWriter, r *http.Request) {
	// Decodificar la película del cuerpo de la solicitud
	var newMovie models.Movie
	err := json.NewDecoder(r.Body).Decode(&newMovie)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para agregar la película a la base de datos
	err = services.AddMovie(mc.DB, &newMovie)
	if err != nil {
		http.Error(w, "Error adding movie to the database", http.StatusInternalServerError)
		return
	}

	// Responder con éxito
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMovie)
}

func (mc *MovieController) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID de la película de los parámetros de la URL
	vars := mux.Vars(r)
	movieID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para eliminar la película
	err = services.DeleteMovie(mc.DB, movieID)
	if err != nil {
		http.Error(w, "Error deleting movie", http.StatusInternalServerError)
		return
	}

	//Responde con exito
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": fmt.Sprintf("Movie with ID %d deleted successfully", movieID),
	}
	json.NewEncoder(w).Encode(response)
}

func (mc *MovieController) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID de la película de los parámetros de la URL
	vars := mux.Vars(r)
	movieID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	// Decodificar la película actualizada del cuerpo de la solicitud
	var updatedMovie models.Movie
	err = json.NewDecoder(r.Body).Decode(&updatedMovie)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para actualizar la película
	err = services.UpdateMovie(mc.DB, movieID, updatedMovie)
	if err != nil {
		http.Error(w, "Error updating movie", http.StatusInternalServerError)
		return
	}

	// Actualizar el ID en la respuesta JSON
	updatedMovie.ID = movieID

	// Responder con éxito
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedMovie)
}

func (mc *MovieController) GetMovies(w http.ResponseWriter, r *http.Request) {
	// Llamar al servicio para obtener la lista de todas las películas
	movies, err := services.GetMovies(mc.DB)
	if err != nil {
		http.Error(w, "Error getting movies", http.StatusInternalServerError)
		return
	}

	// Responder con la lista de películas en formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func (mc *MovieController) GetMovieDetails(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID de la película de los parámetros de la URL
	vars := mux.Vars(r)
	movieID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para obtener los detalles de la película por ID
	movie, err := services.GetMovieDetails(mc.DB, movieID)
	if err != nil {
		http.Error(w, "Error getting movie details", http.StatusInternalServerError)
		return
	}

	// Responder con los detalles de la película en formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func (mc *MovieController) GetMoviesByName(w http.ResponseWriter, r *http.Request) {
	// Obtener el parámetro 'nombre' de la URL
	nombre := r.URL.Query().Get("nombre")

	// Llamar al servicio para obtener la lista de películas con filtro por nombre
	movies, err := services.GetMoviesByName(mc.DB, nombre)
	if err != nil {
		http.Error(w, "Error getting movies", http.StatusInternalServerError)
		return
	}

	// Responder con la lista de películas en formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
