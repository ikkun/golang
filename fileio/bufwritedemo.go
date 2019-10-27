package main

import (
	"bufio"
	"os"
)

func main() {
	f, err := os.Create("newfile.txt")

	if err != nil {
		panic(err)
	}
	buf := bufio.NewWriter(f)
	buf.WriteString("Writing Data to Buffer\n")
	buf.Flush() //Move the buffer data and write to the file
	f.Close()

}
