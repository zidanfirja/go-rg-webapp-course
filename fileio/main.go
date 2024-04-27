package main

import (
	"bufio"
	"fmt"
	"os"
)

// Read and Print File
/*
	1. Open file input.txt
	2. Read content of file
	3. Print content of file
	4. Modify the content of file adding text "Edited"
	5. Save the modified content to output.txt
*/

func main() {
	// open file
	fl, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fl.Close()

	// create output file
	flOutput, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer flOutput.Close()

	// read content of file input.txt
	scanner := bufio.NewScanner(fl)
	for scanner.Scan() {
		// modify the content of file adding text "Edited"
		modifiedText := scanner.Text() + " Edited"
		_, err := flOutput.WriteString(modifiedText + "\n")
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
		fmt.Printf("Original: %s, Modified: %s \n", scanner.Text(), modifiedText)
	}

	// handle error
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}
