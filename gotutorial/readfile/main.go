package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	// b := make([]byte, 1024)
	// n, err := f.Read(b)
	if err != nil {
		log.Fatal(err)
	}

	// b = b[:n]
	fmt.Print(string(b))
}
