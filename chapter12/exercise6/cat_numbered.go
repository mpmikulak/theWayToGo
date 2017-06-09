package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var Number = flag.Bool("n", false, "Adds line numbers to the printout.")

func cat(r *bufio.Reader) {
	lineNumber := 1
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if *Number {
			fmt.Fprintf(os.Stdout, "%v: %s", lineNumber, buf)
			lineNumber++
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
	return
}

func main() {
	flag.Parse()
	flag.PrintDefaults()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s\n", os.Args[0],
				flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}
