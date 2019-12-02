package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("inputData.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		var n int
		n, err = strconv.Atoi(scanner.Text())
		moduleFuel := calculateModuleFuel(n)
		sum = sum + moduleFuel
	}
	println(sum)
}

func calculateModuleFuel(mass int) int {
	dividedMass := mass / 3
	result := dividedMass - 2
	if result <= 0 {
		return 0
	}
	return result + calculateModuleFuel(result)
}
