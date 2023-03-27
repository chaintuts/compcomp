// This file contains code that compares 2 files
//
// Author: Josh McIntyre
//
package main

import (
	"fmt"
	"math"
	"os"
)

// Return the file size in bytes of the specified filename
func file_size(filename string) int64 {

	// Attempt to stat the file and get the size
	stat, err := os.Stat(filename)
	if err != nil {
		fmt.Println(err)
	}

	size := stat.Size()
	
	return size
}

// Return the percentage difference in file size
func size_diff(size1 int64, size2 int64) int64 {

	diff := math.Abs(float64(size1) - float64(size2))
	bigger := math.Max(float64(size1), float64(size2))

	diff_pct := int64((diff / bigger) * 100)

	return diff_pct

}

// The main entry point for the program
func main() {

	// Fetch filenames from command line arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage ./compcomp <filename1> <filename2>")
		return 
	}

	args := os.Args[1:]
	filename1 := args[0]
	filename2 := args[1]
	
	size1 := file_size(filename1)
	size2 := file_size(filename2)
	diff_pct := size_diff(size1, size2)
	
	fmt.Printf("File 1 size (bytes): %d\n", size1)
	fmt.Printf("File 2 size (bytes): %d\n", size2)
	fmt.Printf("File size difference: %d\\%\n", diff_pct)
	
}