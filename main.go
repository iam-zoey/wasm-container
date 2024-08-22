package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Job     string  `json:"job"`
	Address Address `json:"address"`
}

func main() {
	filePath := os.Getenv("JSON_FILE_PATH")
	if filePath == "" {
		log.Fatal("JSON_FILE_PATH environment variable is not set")
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Cannot open the JSON file: %v", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Cannot read the JSON file: %v", err)
	}

	var person Person
	json.Unmarshal(byteValue, &person)
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("Job: %s\n", person.Job)
	fmt.Printf("Address:\n")
	fmt.Printf("  Street: %s\n", person.Address.Street)
	fmt.Printf("  City: %s\n", person.Address.City)
	fmt.Printf("  State: %s\n", person.Address.State)
}
