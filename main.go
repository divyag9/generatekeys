package main

import (
	"log"

	"github.com/divyag9/generatekeys/packages/asymmetric"
)

func main() {
	err := asymmetric.GenerateKeys()
	if err != nil {
		log.Println("Failed to generate keys", err)
	}
	log.Println("Generated keys")
}
