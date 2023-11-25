package main

import "fmt"

func main() {
	resultA := romanToInt("III")
	resultB := romanToInt("LVIII")
	resultC := romanToInt("MCMXCIV")
	fmt.Println(resultA)
	fmt.Println(resultB)
	fmt.Println(resultC)
}

func romanToInt(s string) int {
	symbols := getSymbols()
	cs := []rune(s)
	var sum int
	for i := 0; i < len(s)-1; i++ {
		curr := symbols[string(cs[i])]
		next := symbols[string(cs[i+1])]
		if curr >= next {
			sum += curr
		} else {
			sum -= curr
		}
	}

	return sum + symbols[string(cs[len(cs)-1])]
}

func getSymbols() map[string]int {
	symbols := make(map[string]int)
	symbols["I"] = 1
	symbols["V"] = 5
	symbols["X"] = 10
	symbols["L"] = 50
	symbols["C"] = 100
	symbols["D"] = 500
	symbols["M"] = 1000

	return symbols
}
