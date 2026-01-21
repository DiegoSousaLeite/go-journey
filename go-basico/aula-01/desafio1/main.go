package main

import ("fmt"; "time")

func main() {

	var age int
	fmt.Println("Please enter your age:")
	fmt.Scanln(&age)

	now := time.Now()
	year := now.Year() - age

	fmt.Printf("You were born in the year %d.\n", year)
}
