package main

import "fmt"

func Count(chars string) int {
	return len(chars)
}

func main() {
	var toBeCounted string
	fmt.Print("What is the input string? ")
	_, err := fmt.Scanln(&toBeCounted)
	if err != nil {
		fmt.Println("Couldn't get input string")
	}
	totalChars := Count(toBeCounted)
	fmt.Printf("%v has %v characters.\n", toBeCounted, totalChars)

}
