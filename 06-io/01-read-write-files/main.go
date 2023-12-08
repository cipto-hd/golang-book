package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var path = "invoices.csv"
	filebuffer, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}
	var inputdata string = string(filebuffer)

	total := 0
	rows := strings.Split(inputdata, "\n")
	for _, row := range rows {
		fmt.Println("row:", row)
		cols := strings.Split(row, ",")

		// challenge: calculate the total
		col2 := strings.TrimSpace(cols[1])
		if col2 != "amount" {
			x, _ := strconv.Atoi(col2)
			total += x
		}
	}
	fmt.Println("The total is ", total)
}
