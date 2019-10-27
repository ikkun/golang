package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("myfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 30)
	f.Seek(5, 0) //seek the file pointer to 5th
	f.Read(data)
	fmt.Printf("The file data is %s\n", string(data))
	defer f.Close()
}
