package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"adventofcode"
)

const (
	id = "0102"
)

func lineToInt(line string) int {
	i, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	return i
}

func scannerToIntSlice(scanner *bufio.Scanner) []int {
	var numbers []int

	for scanner.Scan() {
		numbers = append(numbers, lineToInt(scanner.Text()))
	}

	return numbers
}

func sum3(start int, numbers []int) (int, error) {
	sum := 0
	for i := start; i < start+3; i++ {
		if start > len(numbers)-3 {
			return 0, io.EOF
		}

		sum += numbers[i]
	}

	return sum, nil
}

func main() {
	input := adventofcode.NewInput[int](id+".input", func(line string) int {
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		return i
	})

	data := input.Data()
	if len(data) == 0 {
		panic("do input data")
	}

	previous, err := sum3(0, data)
	if err != nil {
		panic(err)
	}

	increases := 0
	for i := 1; i < len(data); i++ {
		current, err := sum3(i, data)
		if err == io.EOF {
			break
		}

		if current > previous {
			increases++
		}

		previous = current
	}

	fmt.Println(increases)
}
