package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func checkBoard(board [][]int, nums []int) (win bool, unmarked int) {
	win = false
	unmarked = 0

	//check rows
	for _, row := range board {
		hits := 0
		for _, v := range row {
			found := false
			for _, n := range nums {
				if n == v {
					hits++
					found = true
					// fmt.Println("found", v, hits)
					break
				}
			}
			if !found {
				unmarked += v
			}
		}
		if hits == 5 {
			win = true
		}
	}

	return
}

func transpose(slice [][]int) [][]int {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]int, xl)
	for i := range result {
		result[i] = make([]int, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func partA() {
	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")

	n := strings.Split(lines[0], ",")
	nums := make([]int, len(n))
	for i, v := range n {
		nums[i], _ = strconv.Atoi(v)
	}
	fmt.Println(nums)

	boards := make([][][]int, 0)
	for i := 2; i < len(lines); i += 6 {
		board := make([][]int, 0)
		for j := 0; j < 5; j++ {
			boardline := make([]int, 5)
			s := strings.Fields(lines[i+j])
			for k := 0; k < 5; k++ {
				boardline[k], _ = strconv.Atoi(s[k])
			}
			board = append(board, boardline)
		}
		boards = append(boards, board)
	}

	transposedBoards := make([][][]int, len(boards))
	for i, board := range boards {
		transposedBoards[i] = transpose(board)
	}

	fmt.Println(boards[0])
	fmt.Println(transposedBoards[0])

	win := false
	sum := 0
	wonBoards := make([]int, 0)
	for n := 5; n < len(nums); n++ {
		//fmt.Println("checking ", nums[:n])
		for i := range boards {
			skip := false
			for _, b := range wonBoards {
				if b == i {
					skip = true
				}
			}
			if skip {
				continue
			}
			win, sum = checkBoard(boards[i], nums[:n])
			if win {
				fmt.Printf("board %d won\n", i)
				fmt.Println(n, sum, nums[n-1], sum*nums[n-1])
				wonBoards = append(wonBoards, i)
			}

			win, sum = checkBoard(transposedBoards[i], nums[:n])

			if win {
				fmt.Printf("board %d won\n", i)
				fmt.Println(n, sum, nums[n-1], sum*nums[n-1])
				wonBoards = append(wonBoards, i)
			}
			//fmt.Println(board)
		}
	}

	//fmt.Println(boards)
}

func partB() {
}

func main() {
	partA()
	partB()
}
