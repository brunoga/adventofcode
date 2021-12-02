package main

import (
	"fmt"
	"strconv"

	"adventofcode"
)

const (
	id = "0101"
)

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

	previous := data[0]

	increases := 0
	for i := 1; i < len(data); i++ {
		current := data[i]

		if current > previous {
			increases++
		}

		previous = current
	}

	fmt.Println(increases)
}

