package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func partA() {
	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")

	position := 0
	depth := 0

	for _, l := range lines {

		l = strings.TrimSpace(l)
		instruction := strings.Split(l, " ")
		direction := instruction[0]
		value, _ := strconv.Atoi(instruction[1])
		switch direction {
		case "forward":
			position += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}

	fmt.Println(position, depth, position*depth)
}

func partB() {
	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")

	position := 0
	aim := 0
	depth := 0

	for _, l := range lines {

		l = strings.TrimSpace(l)
		instruction := strings.Split(l, " ")
		direction := instruction[0]
		value, _ := strconv.Atoi(instruction[1])
		switch direction {
		case "forward":
			position += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}

	fmt.Println(position, depth, position*depth)
}

func main() {
	partA()
	partB()
}
