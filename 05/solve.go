package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func printBoard(board [][]int) {
	for _, r := range board {
		for _, c := range r {
			if c == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(c)
			}
		}
		fmt.Println()
	}
}

func partA() {
	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	maxX := 0
	maxY := 0

	lineCoordinates := make([]line, 0)
	for _, l := range lines {
		coords := strings.Split(l, " -> ")
		start := strings.Split(coords[0], ",")
		end := strings.Split(coords[1], ",")
		x1, _ := strconv.Atoi(start[0])
		y1, _ := strconv.Atoi(start[1])
		x2, _ := strconv.Atoi(end[0])
		y2, _ := strconv.Atoi(end[1])
		//if x1 == x2 || y1 == y2 {
		if x1 > maxX {
			maxX = x1
		}
		if x2 > maxX {
			maxX = x2
		}
		if y1 > maxY {
			maxY = y1
		}
		if y2 > maxY {
			maxY = y2
		}
		lineCoordinates = append(lineCoordinates, line{x1, y1, x2, y2})
		//}
	}

	board := make([][]int, maxY+1)
	for i := range board {
		board[i] = make([]int, maxX+1)
	}

	for _, coords := range lineCoordinates {
		// printBoard(board)
		// fmt.Println(coords)
		if coords.x1 == coords.x2 {
			if coords.y1 < coords.y2 {
				for y := coords.y1; y <= coords.y2; y++ {
					board[y][coords.x1] += 1
				}
			}
			if coords.y2 <= coords.y1 {
				for y := coords.y2; y <= coords.y1; y++ {
					board[y][coords.x1] += 1
				}
			}
		} else if coords.y1 == coords.y2 {
			//horizontal
			if coords.x1 < coords.x2 {
				for x := coords.x1; x <= coords.x2; x++ {
					board[coords.y1][x] += 1
				}
			}
			if coords.x2 <= coords.x1 {
				for x := coords.x2; x <= coords.x1; x++ {
					board[coords.y1][x] += 1
				}
			}

		} else {
			xDir := 1
			yDir := 1
			length := int(math.Abs(float64(coords.x2 - coords.x1)))
			if coords.x1 > coords.x2 {
				xDir = -1
			}
			if coords.y1 > coords.y2 {
				yDir = -1
			}
			for i := 0; i <= length; i++ {
				board[coords.y1+yDir*i][coords.x1+xDir*i] += 1
			}
		}
	}
	printBoard(board)
	fmt.Println(lineCoordinates)
	count := 0
	for _, r := range board {
		for _, c := range r {
			if c > 1 {
				count++
			}
		}
	}
	fmt.Println(count)
}

func partB() {
}

func main() {
	partA()
	partB()
}
