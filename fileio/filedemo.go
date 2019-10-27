package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("myfile.txt") //Create a new file

	if err != nil {
		log.Fatal(err)
	}

	data := []byte("Hello world\n")
	f.Write(data)                     //Data as a byte array
	f.WriteString("This is a String") //data as string
	defer f.Close()
}
