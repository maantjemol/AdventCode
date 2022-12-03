package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func alph_score(l string) int{
	alph_array := [52]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z" }
	for i := 0; i < len(alph_array); i++ {
		if l == alph_array[i]{
			return i + 1
		}
	}
	return 0
}


func main() {
	input, err := os.ReadFile("./input.txt")
    check(err)

	split_input := strings.Split(string(input), "\n")

	
	// []
	// 

	total_score := 0

	for i := 0; i < len(split_input); i+=3 {
		sack_1 := split_input[i]
		sack_2 := split_input[i+1]
		sack_3 := split_input[i+2]

		
		
		for j := 0; j < len(sack_1); j++ {
			if strings.Contains(sack_2, string(sack_1[j])) && strings.Contains(sack_3, string(sack_1[j])) {
				total_score += alph_score(string(sack_1[j]))
				break
			} else {
				
			}
		}
	}
	
	fmt.Println(total_score)

	// fmt.Println(split_input)
}
