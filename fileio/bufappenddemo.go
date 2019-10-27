package main

import (
	"os"
)

func main() {
	f, _ := os.OpenFile("newfile.txt", os.O_APPEND|os.O_WRONLY, 0600)
	f.WriteString("test append")

	f.Close()

}
