package services

import (
	"database/sql"
	"errors"
	"fmt"
	models "labora/clases/Practicas/practica14_9-01-2023/movie-api/models"
)

func AddMovie(db *sql.DB, movie *models.Movie) error {
	// Lógica para insertar la película en la base de datos
	// Por ejemplo:
	query := "INSERT INTO movie (name, created, stock, price) VALUES ($1, $2, $3, $4) RETURNING id"
	err := db.QueryRow(query, movie.Name, movie.Created, movie.Stock, movie.Price).Scan(&movie.ID)
	if err != nil {
		fmt.Println("Error inserting movie into database:", err)
		return fmt.Errorf("error inserting movie into database: %w", err)
	}

	return nil
}

func DeleteMovie(db *sql.DB, movieID int) error {
	// Preparar la consulta SQL para eliminar la película por ID
	query := "DELETE FROM movie WHERE id = $1"
	_, err := db.Exec(query, movieID)
	if err != nil {
		return fmt.Errorf("error deleting movie from database: %w", err)
	}

	return nil
}

func UpdateMovie(db *sql.DB, movieID int, updatedMovie models.Movie) error {
	// Preparar la consulta SQL para actualizar la película por ID
	query := "UPDATE movie SET name = $2, created = $3, stock = $4, price = $5 WHERE id = $1"
	_, err := db.Exec(query, movieID, updatedMovie.Name, updatedMovie.Created, updatedMovie.Stock, updatedMovie.Price)
	if err != nil {
		return fmt.Errorf("error updating movie in database: %w", err)
	}

	return nil
}

func GetMovies(db *sql.DB) ([]models.Movie, error) {
	// Lógica para obtener la lista de todas las películas
	// Por ejemplo:
	query := "SELECT id, name, created, stock, price FROM movie"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error getting movies: %w", err)
	}
	defer rows.Close()

	var movies []models.Movie

	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.ID, &movie.Name, &movie.Created, &movie.Stock, &movie.Price)
		if err != nil {
			return nil, fmt.Errorf("error scanning movie row: %w", err)
		}
		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over movie rows: %w", err)
	}

	return movies, nil
}

func GetMovieDetails(db *sql.DB, movieID int) (*models.Movie, error) {
	// Lógica para obtener los detalles de la película por ID
	// Por ejemplo:
	query := "SELECT id, name, created, stock, price FROM movie WHERE id = $1"
	row := db.QueryRow(query, movieID)

	// Crear una estructura para almacenar los detalles de la película
	var movie models.Movie
	err := row.Scan(&movie.ID, &movie.Name, &movie.Created, &movie.Stock, &movie.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("movie with ID %d not found", movieID)
		}
		return nil, fmt.Errorf("error getting movie details: %w", err)
	}

	return &movie, nil
}

func GetMoviesByName(db *sql.DB, nombre string) ([]models.Movie, error) {
	// Lógica para obtener la lista de películas con filtro por nombre
	// Por ejemplo:
	query := "SELECT id, name, created, stock, price FROM movie WHERE name ILIKE '%' || $1 || '%'"
	rows, err := db.Query(query, nombre)
	if err != nil {
		return nil, fmt.Errorf("error getting movies by name: %w", err)
	}
	defer rows.Close()

	var movies []models.Movie

	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.ID, &movie.Name, &movie.Created, &movie.Stock, &movie.Price)
		if err != nil {
			return nil, fmt.Errorf("error scanning movie row: %w", err)
		}
		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over movie rows: %w", err)
	}

	return movies, nil
}
