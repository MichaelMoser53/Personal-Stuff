package main

import (
		"fmt"
		
				)

// sameWords array declared and set to words array and printed
func main() {
	
    fmt.Print("Welcome to Golang\n")
    words := []string {"Hello", "This", "Is", "An", "Array"}
    sameWords := []string{}
    sameWords = words
    
    for i := 0; i < len(sameWords); i++{
    	fmt.Print(sameWords[i], " ")
    }

        
    numbers := []int {2,37,45,20000000}
    fmt.Print("\n")
    for j := 0; j < len(numbers); j++{
    	
    	fmt.Print(numbers[j], " ")
    }

    var sumation int = numbers[0] + numbers[2]
    fmt.Println("\nThis is the sum of the 1st and 3rd element of the numbers array:", sumation)

var num int
for x := 0; x < 10; x++{
	num = num + 10
	fmt.Print(num, " ")
}

	
}
