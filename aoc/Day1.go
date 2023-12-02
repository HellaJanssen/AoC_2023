package aoc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

type Day1 struct{}

func (d Day1) Name() string {
	return "Day 1"
}

func firstDigitInString(line string) (int, error) {
	for len(line) > 0 {
		char, size := utf8.DecodeRuneInString(line)
		line = line[size:]
		if !unicode.IsDigit(char) {
			continue
		}

		digit, _ := strconv.Atoi(string(char))
		return digit, nil
	}

	return 0, fmt.Errorf("No digit in string")
}

func lastDigitInString(line string) (int, error) {
	for len(line) > 0 {
		char, size := utf8.DecodeLastRuneInString(line)
		if unicode.IsDigit(char) {
			digit, _ := strconv.Atoi(string(char))
			return digit, nil
		}
		line = line[:len(line)-size]
	}

	return 0, fmt.Errorf("No digit string")
}

func (d Day1) Solution() string {
	file, error := os.Open("aoc/Day1Input.txt")

	if error != nil {
		log.Fatal(error)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int

	for scanner.Scan() {
		line := scanner.Text()

		a, _ := firstDigitInString(line)
		b, _ := lastDigitInString(line)

		sum += (a * 10) + b
	}

	return strconv.FormatInt(int64(sum), 10)
}
