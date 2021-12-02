package adventofcode

import (
	"bufio"
	"os"
)

type Input[T any] struct {
	data []T
}

func NewInput[T any](inputFile string, lineParser func(string) T) *Input[T] {
	f, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	data := []T{}

	for scanner.Scan() {
		data = append(data, lineParser(scanner.Text()))
	}

	return &Input[T]{
		data: data,
	}
}

func (i *Input[T]) Data() []T {
	return i.data
}
