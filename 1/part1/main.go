package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int

	for scanner.Scan() {
		word := scanner.Text()
		value := extractNumbers(word)

		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Error converting to integer:", err)
			continue
		}

		sum += num
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Sum of all calibration values:", sum)

}

func extractNumbers(word string) string {
	var first, last rune
	for _, char := range word {
		if unicode.IsDigit(char) {
			if first == 0 {
				first = char
			}
			last = char
		}
	}

	return fmt.Sprintf("%s%s", string(first), string(last))
}
