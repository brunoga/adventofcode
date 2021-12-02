package main

import (
	"adventofcode"
	"fmt"
	"strconv"
	"strings"
)

const (
	id = "0201"
)

type Movement struct {
	direction string
	distance  int
}

func main() {
	input := adventofcode.NewInput[Movement](id+".input", func(line string) Movement {
		fields := strings.Split(line, " ")

		i, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}

		return Movement{
			fields[0],
			i,
		}
	})

	data := input.Data()
	if len(data) == 0 {
		panic("no input data")
	}

	x := 0
	y := 0

	for _, movement := range data {
		switch movement.direction {
		case "forward":
			x += movement.distance
			break
		case "backward":
			x -= movement.distance
			break
		case "up":
			y -= movement.distance
			break
		case "down":
			y += movement.distance
			break
		default:
			panic("unexpected direction")
		}
	}

	fmt.Println(x * y)
}
