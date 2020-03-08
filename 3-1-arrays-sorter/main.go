package main

import (
	"strconv"
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Number Sorter
//
//  Your goal is to sort the given numbers from the command-line.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Sort the given numbers and print them.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments are not a valid number, skip it
//
// HINTS
//  + You can use the bubble-sort algorithm to sort the numbers.
//    Please watch this: https://youtu.be/nmhjrI-aW5o?t=7
//
//  + When swapping the elements, do not check for the last element.
//
//    Or, you will receive this error:
//    "panic: runtime error: index out of range"
//
// EXPECTED OUTPUT
//   go run main.go
//     Please give me up to 5 numbers.
//
//   go run main.go 6 5 4 3 2 1
//     Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...
//
//   go run main.go 5 4 3 2 1
//     [1 2 3 4 5]
//
//   go run main.go 5 4 a c 1
//     [0 0 1 4 5]
// ---------------------------------------------------------

func main() {
	const (
		usage = "Please give me up to 5 numbers."
	)

	args := os.Args[1:]

	var (
		array [5]float64
		value float64
		err error
	)

	if len(args) == 0 || len(args) > 5 {
		fmt.Println(usage)
		return
	}

	for i, v := range args {
		if value, err = strconv.ParseFloat(v, 64); err != nil {
			continue
		}
		array[i] = value
	}
	
	fmt.Printf("%v\n", array)

	for range array {
		for i, v := range array {
			if i < len(array)-1 && v > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
			}
		}
	}

	fmt.Printf("%v", array)
}