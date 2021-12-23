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

	days := make([]int, 9)
	for _, l := range lines {
		ages := strings.Split(l, ",")
		for _, s := range ages {
			age, _ := strconv.Atoi(s)
			days[age] = days[age] + 1
		}
	}

	fmt.Println("initial:", days)
	for d := 0; d < 256; d++ {
		fmt.Println("day ", d+1)
		//give birth
		newborns := days[0]

		//age
		for i, c := range days {
			if i > 0 {
				days[i-1] = c
			}
		}
		days[6] += newborns
		days[8] = newborns

		fmt.Println("giving birth:", newborns)
		fmt.Println(days)

		count := 0
		for _, value := range days {
			count += value
		}
		fmt.Println("population:", count)
	}

	fmt.Println(days)
	count := 0
	for _, value := range days {
		count += value
	}
	fmt.Println(count)
}

func partB() {
}

func main() {
	partA()
	partB()
}
