package main

import (
	"class1"
	"log"
	"os"
)

func main() {
	add := class1.Add(2, 3)
	if add != 5 {
		log.Printf("Error: want %d, got %d", 5, add)
		os.Exit(1) // whatever error code not equal 0 is an error
	}
	log.Println("All is ok")
}
