package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./day_1.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Day 1.1: %v", SolveDay11(data))
}

func SolveDay11(data []byte) int {
	sum := 0

	lines := bytes.Split(data, []byte("\n"))

	for _, line := range lines {
		sum += FirstLastDigitNum(line)
	}
	return sum
}

func FirstLastDigitNum(line []byte) int {
	first_digit := func() int {
		for _, v := range line {
			if num := v - 48; num <= 9 {
				return int(num)
			}
		}
		return 0
	}()
	last_digit := func() int {
		for index := range line {
			if num := line[len(line)-index-1] - 48; num <= 9 {
				return int(num)
			}
		}
		return 0
	}()

	return 10*first_digit + last_digit
}
