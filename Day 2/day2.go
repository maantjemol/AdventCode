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


func main() {
	shapes := make(map[string]int)

	shapes["A"] = 1 // Rock
	shapes["B"] = 2 // Paper
	shapes["C"] = 3 // Scissors

	options := [3]string{"A", "B", "C"}

	// shape me - shape op = 1 == win

	shapes["X"] = 1 // Lose
	shapes["Y"] = 2 // Draw
	shapes["Z"] = 3 // Win

	input, err := os.ReadFile("./input_2.txt")
    check(err)
	split_input := strings.Split(string(input), "\n")

	total_score := 0
	
	for i := 0; i < len(split_input); i++ {
		choices := strings.Split(split_input[i], " ")
		op := choices[0]
		outcome := choices[1]

		fmt.Println("\nRound:", i)
		
		// Draw
		if outcome == "Y"{
			total_score += 3 + shapes[op]
		}

		// Win
		if outcome == "Z"{
			winIndex := shapes[op]
			if winIndex == 3{
				winIndex = 0
			}

			total_score += 6 + winIndex + 1
		}

		// Lose
		if outcome == "X"{
			winIndex := shapes[op] - 2
			if winIndex == -1{
				winIndex = 2
			}

			total_score += shapes[options[winIndex]]
		}

		fmt.Println(total_score)
	}

	fmt.Println(total_score)


	// fmt.Println(split_input)
}
