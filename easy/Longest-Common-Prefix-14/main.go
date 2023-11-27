package main

import (
	"fmt"
)

func main() {
	strsA := []string{"flower", "flow", "flight"}
	resultA := longestCommonPrefix(strsA)
	fmt.Println(resultA)

	strsB := []string{"dog", "racecar", "car"}
	resultB := longestCommonPrefix(strsB)
	fmt.Println(resultB)

	strsC := []string{"ab", "a"}
	resultC := longestCommonPrefix(strsC)
	fmt.Println(resultC)
}

func longestCommonPrefix(strs []string) string {
	var data [][]rune
	for _, str := range strs {
		data = append(data, []rune(str))
	}

	var count int
	var isDiffer bool
	for indexOfPointer := 0; indexOfPointer < len(data[0]); indexOfPointer++ {
		for indexOfElement := 1; indexOfElement < len(strs); indexOfElement++ {
			isOutOfBound := indexOfPointer >= len(strs[indexOfElement])
			if isOutOfBound || data[0][indexOfPointer] != data[indexOfElement][indexOfPointer] {
				isDiffer = true
				break
			}
		}

		if isDiffer {
			break
		}

		count++
	}

	if count == 0 {
		return ""
	} else {
		return string(data[0][:count])
	}
}
