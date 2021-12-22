package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	var nums = make([]int, 0, len(lines))

	for _, l := range lines {
		n, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}

	count := 0
	last := -1
	for pos, depth := range nums {
		if pos != 0 && depth > last {
			count++
		}
		last = depth
	}
	fmt.Println(count)

	windowSums := make([]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		window := nums[i : i+3]
		//fmt.Println(window)
		sum := 0
		for _, v := range window {
			sum += v
		}
		windowSums = append(windowSums, sum)
	}

	//fmt.Println(windowSums)

	count = 0
	last = -1
	for pos, depth := range windowSums {
		if pos != 0 && depth > last {
			count++
		}
		last = depth
	}
	fmt.Println(count)

}
