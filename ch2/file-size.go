package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("go.mod")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size(), "bytes")
}
