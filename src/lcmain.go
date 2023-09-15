package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("problems.json", os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		log.Fatal("Error opening file")
	}

	defer file.Close()
	/* cli parsing here? */
	read(file)
	fmt.Println(lcproblems)
	write(file)
}
