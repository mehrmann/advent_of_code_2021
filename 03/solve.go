package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func partA(bits []string) {
	s := mostCommonBits(bits)
	fmt.Println(s)

	gamma := 0
	epsilon := 0
	for i := 0; i < len(s); i++ {
		b := 0
		if s[len(s)-i-1] == '1' {
			b = 1
		}
		gamma = gamma | b<<i
		epsilon = epsilon | int(math.Abs(float64(b-1)))<<i
	}

	fmt.Println(gamma, epsilon, gamma*epsilon)
}

func mostCommonBits(bits []string) string {
	s := ""
	for i := 0; i < len(bits[0]); i++ {
		zeroes := 0
		ones := 0
		for _, s := range bits {
			switch s[i] {
			case '0':
				zeroes++
			case '1':
				ones++
			}
		}
		if zeroes > ones {
			s = s + "0"
		} else {
			s = s + "1"
		}
	}
	return s
}

func bitsToDec(s string) int {
	r := 0
	for i := 0; i < len(s); i++ {
		b := 0
		if s[len(s)-i-1] == '1' {
			b = 1
		}
		r |= b << i
	}
	return r
}

func filter(bits []string, filter string, pos int) []string {
	var newbits = make([]string, 0)
	for _, b := range bits {
		if b[pos] == filter[pos] {
			newbits = append(newbits, b)
		}
	}
	return newbits
}

func inverse(s string) string {
	r := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			r += "0"
		} else {
			r += "1"
		}
	}
	return r
}

func partB(bits []string) {
	common := mostCommonBits(bits)
	least := inverse(common)
	fmt.Println(bits)
	fmt.Println(common)
	fmt.Println(least)
	oxygen := bits
	co2 := bits
	for i := 0; i < len(bits); i++ {
		oxygen = filter(oxygen, common, i)
		fmt.Println(oxygen)
		common = mostCommonBits(oxygen)
		if len(oxygen) == 1 {
			break
		}
	}
	for i := 0; i < len(bits); i++ {
		co2 = filter(co2, least, i)
		least = inverse(mostCommonBits(co2))
		fmt.Println(co2)
		if len(co2) == 1 {
			break
		}
	}
	fmt.Println(oxygen[0], co2[0])
	fmt.Println(bitsToDec(oxygen[0]), bitsToDec(co2[0]))
	fmt.Println(bitsToDec(oxygen[0]) * bitsToDec(co2[0]))
}

func main() {
	b, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	var bits = make([]string, 0, len(lines))

	for _, l := range lines {
		bits = append(bits, l)
	}

	partA(bits)

	partB(bits)

}
