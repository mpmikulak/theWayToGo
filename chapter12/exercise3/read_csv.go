package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	Title    string
	Price    float64
	Quantity int
}

func main() {
	// Open the file
	f, err := os.Open("products.txt")
	defer f.Close() // Defer closing the file
	if err != nil { // Catch errors
		return
	}

	rec := make([]Product, 0)

	// Convert the *File to a reader
	fr := bufio.NewReader(f)
	for {
		// Read a line from the file
		line, err := fr.ReadString('\n')
		if err == io.EOF {
			break
		}

		// Split the line by semi-colon
		parts := strings.Split(line, ";")
		for ix := range parts {
			parts[ix] = strings.Trim(parts[ix], "\n")
		}

		entry := new(Product)
		entry.Title = parts[0]
		entry.Price, err = strconv.ParseFloat(parts[1], 64)
		if err != nil {
			fmt.Println("OH NOES1")
		}
		entry.Quantity, err = strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("OH NOES2")
		}
		rec = append(rec, *entry)
	}
	for ix, _ := range rec {
		fmt.Printf("There are %v of %v and they cost $%v.\n", rec[ix].Quantity, rec[ix].Title, rec[ix].Price)
	}
}
