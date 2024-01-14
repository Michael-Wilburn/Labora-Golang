package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Usuario struct {
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad"`
}

func main() {
	user1 := Usuario{"Luuis", 8}
	userJSON, err := json.Marshal(user1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(userJSON))
	fmt.Println(bytes.NewBuffer(userJSON))
}
