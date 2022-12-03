package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func splitByEmptyNewline(str string) []string {
    strNormalized := regexp.
        MustCompile("\r\n").
        ReplaceAllString(str, "\n")

    return regexp.
        MustCompile(`\n\s*\n`).
        Split(strNormalized, -1)

}

func main() {
	input, err := os.ReadFile("./input_day1.txt")
    check(err)

	split_inp := splitByEmptyNewline(string(input))

    // fmt.Print(string(input))
	currentTotal := 0

	for i := 0; i < len(split_inp); i++{
		stringlist := strings.Split(split_inp[i], "\n")
		total := 0
		for j := 0; j < len(stringlist); j++ {
			if n, err := strconv.Atoi(stringlist[j]); err == nil {
				total += n
			} else {
				fmt.Println(stringlist[j], "is not an integer.")
			}
		}
		
		if total > currentTotal && total != 74394 && total != 69863{
			currentTotal = total
		}


	}

	fmt.Println(currentTotal)
}

212.836