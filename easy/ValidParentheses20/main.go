package main

import "fmt"

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid("(){}}{"))
}

func isValid(s string) bool {
	leftSide := getLeftSideMap()
	rightSide := getRightSideMap()

	if len(s) < 2 {
		return false
	}

	var store []string
	for _, c := range []rune(s) {
		current := string(c)
		isLeftSide := leftSide[current] != ""
		if isLeftSide {
			store = append(store, current)
		} else {
			if len(store) < 1 || (rightSide[current] != store[len(store)-1]) {
				return false
			}
			store = store[:len(store)-1]
		}
	}

	return len(store) == 0
}

func getLeftSideMap() map[string]string {
	m := map[string]string{}
	m["["] = "]"
	m["{"] = "}"
	m["("] = ")"

	return m
}

func getRightSideMap() map[string]string {
	m := map[string]string{}
	m["]"] = "["
	m["}"] = "{"
	m[")"] = "("

	return m
}
