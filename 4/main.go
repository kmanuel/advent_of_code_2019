package main

import (
	"fmt"
	"strconv"
)

func main() {
	found := 0
	for candidate := 278384; candidate <= 824795; candidate++ {
		if isPossiblePassword(candidate) {
			found++
		}
	}
	fmt.Printf("found=%d", found)
}

func isPossiblePassword(candidate int) bool {
	numAsString := strconv.Itoa(candidate)
	numAsBytes := []byte(numAsString)
	var prev byte
	hasDoubleNumber := false
	doubleCheckPassed := false
	for i := 0; i < len(numAsBytes); i++ {
		curr := numAsBytes[i]
		if i != 0 {
			if !doubleCheckPassed {
				if prev == curr {
					hasDoubleNumber = true
					if i > 1 {
						if numAsBytes[i-2] == curr {
							hasDoubleNumber = false
						}
					}
				} else {
					if hasDoubleNumber {
						doubleCheckPassed = true
					}
				}
			}
			if prev > curr {
				return false
			}
		}
		prev = curr
	}
	return hasDoubleNumber
}
