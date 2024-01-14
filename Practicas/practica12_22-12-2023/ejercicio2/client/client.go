package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

const (
	incrementURL = "http://localhost:9090/api/v1/inc"
	decrementURL = "http://localhost:9090/api/v1/dec"
	currentURL   = "http://localhost:9090/api/v1/curr"
)

type Response struct {
	Value int `json:"value"`
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go sendRequests(incrementURL, 10000, &wg)
	go sendRequests(decrementURL, 1000, &wg)

	wg.Wait()

	getCurrentValue()
}

func sendRequests(url string, count int, wg *sync.WaitGroup) {
	defer wg.Done()

	client := http.Client{}

	for i := 0; i < count; i++ {
		req, err := http.NewRequest(http.MethodPost, url, nil)
		if err != nil {
			fmt.Printf("Error creando la solicitud: %v\n", err)
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error al enviar la solicitud: %v\n", err)
			return
		}

		resp.Body.Close()
	}
}

func getCurrentValue() {
	resp, err := http.Get(currentURL)
	if err != nil {
		fmt.Printf("Error al obtener el valor actual: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result Response                               // Estructura para almacenar la respuesta JSON
		err := json.NewDecoder(resp.Body).Decode(&result) // Decodificar la respuesta JSON
		if err != nil {
			fmt.Printf("Error al decodificar la respuesta JSON: %v\n", err)
			return
		}
		fmt.Printf("Valor actual de la variable: %d\n", result.Value)
	} else {
		fmt.Println("Error al obtener el valor actual:", resp.Status)
	}
}
