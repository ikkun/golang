package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Customers struct {
}

func main() {
	f, err := os.Open("breach.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range records {
		fmt.Println(record[1])
	}
	// reader := csv.NewReader(f)
	// for {
	// 	rows, err := reader.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println(rows)
	// }

}
