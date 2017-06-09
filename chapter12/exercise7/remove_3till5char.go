package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, _ := os.Open("goprogram")
	outputFile, _ := os.OpenFile("goprogramT", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	defer outputWriter.Flush()

	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			fmt.Println("EOF")
			return
		}

		outputString := string([]byte(inputString)[2:5]) + "\n"
		fmt.Println(outputString)
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Fprintln(outputFile, string([]byte(inputString)[2:5]))
	}
	fmt.Println("Conversion Done")
}
