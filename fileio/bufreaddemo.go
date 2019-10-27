package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("newfile.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)
	data, _ := reader.Peek(5)
	// data.Peek(5)
	fmt.Println(string(data))
	f.Close()
}
