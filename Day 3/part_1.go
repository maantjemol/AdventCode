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

	for i := 0; i < len(split_input); i++ {
		sack := split_input[i]
		len_inp := len(sack) / 2
		sack_1 := sack[:len_inp]
		sack_2 := sack[len_inp:]
		fmt.Println(sack, sack_1, sack_2)
		
		for j := 0; j < len_inp; j++ {
			if strings.Contains(sack_1, string(sack_2[j])) {
				total_score += alph_score(string(sack_2[j]))
				fmt.Println("score:", alph_score(string(sack_2[j])), "letter: ", string(sack_2[j]))
				break
			} else {
				
			}
		}
	}
	
	fmt.Println(total_score)

	// fmt.Println(split_input)
}
