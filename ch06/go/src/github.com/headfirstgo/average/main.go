// average calculates the average of several numbers.
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	path, err := os.Getwd()
	fmt.Println("path", path)
	numbers, err := datafile.GetFloats("ch06/go/bin/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
