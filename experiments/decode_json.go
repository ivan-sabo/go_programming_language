package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type People struct {
	Persons []Person `json:"persons"`
}

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	var people People

	data, err := ioutil.ReadFile("data.json")

	err = json.Unmarshal(data, &people)
	if err != nil {
		log.Fatalf("an error occured: %e", err)
	}

	fmt.Printf("decoded: %v", people)
}
