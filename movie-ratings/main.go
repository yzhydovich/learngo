package main

import (
	"strconv"
	"fmt"
	"os"
)

// ---------------------------------------------------------
// STORY
//
//  Your boss wants you to create a program that will check
//  whether a person can watch a particular movie or not.
//
//  Imagine that another program provides the age to your
//  program. Depending on what you return, the other program
//  will issue the tickets to the person automatically.
//
// EXERCISE: Movie Ratings
//
//  1. Get the age from the command-line.
//
//  2. Return one of the following messages if the age is:
//     -> Above 17         : "R-Rated"
//     -> Between 13 and 17: "PG-13"
//     -> Below 13         : "PG-Rated"
//
// RESTRICTIONS:
//  1. If age data is wrong or absent let the user know.
//  2. Do not accept negative age.
//
// BONUS:
//  1. BONUS: Use if statements only twice throughout your program.
//  2. BONUS: Use an if statement only once.
//
// EXPECTED OUTPUT
//  go run main.go 18
//    R-Rated
//
//  go run main.go 17
//    PG-13
//
//  go run main.go 12
//    PG-Rated
//
//  go run main.go
//    Requires age
//
//  go run main.go -5
//    Wrong age: "-5"
// ---------------------------------------------------------

func main() {
	const (
		r = "R-Rated"
		pg13 = "PG-13"
		pg = "PG-Rated"
	)

	if args := os.Args; len(args) != 2 {
		fmt.Println("Error: Requires age")
	} else if age, err := strconv.Atoi(args[1]); err != nil || age < 0 {
		fmt.Printf("Error: wrong age: %q", args[1])
	} else if age < 13 {
		fmt.Println(pg)
	} else if  age >= 13  && age <= 17 {
		fmt.Println(pg13)
	} else {
		fmt.Println(r)
	}
}