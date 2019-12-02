package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	// star1()
	star2()
}

func star1() {
	f := openFile()
	defer f.Close()
	numbers := parse(f)
	result := calculateResult(numbers)
	println(result)
}

func star2() {
	f := openFile()
	defer f.Close()
	numbers := parse(f)
	inputCopy := make([]int, len(numbers))
	copy(inputCopy, numbers)
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			inputCopy[1] = x
			inputCopy[2] = y
			result := calculateResult(inputCopy)
			if result == 19690720 {
				endRes := 100*x + y
				println(endRes)
			}
			copy(inputCopy, numbers)
		}
	}
	print("found none")
}

func calculateResult(memory []int) int {
	for i := 0; i < len(memory); i += 4 {
		instructionCode := memory[i]
		if instructionCode != 1 && instructionCode != 2 && instructionCode != 99 {
			return -1
		}
		firstArgPos := memory[i+1]
		if firstArgPos > len(memory) {
			return -1
		}
		secondArgPos := memory[i+2]
		if secondArgPos > len(memory) {
			return -1
		}
		resultPos := memory[i+3]
		if resultPos > len(memory) {
			return -1
		}

		if instructionCode == 1 {
			memory[resultPos] = memory[firstArgPos] + memory[secondArgPos]
		}
		if instructionCode == 2 {
			memory[resultPos] = memory[firstArgPos] * memory[secondArgPos]
		}
		if instructionCode == 99 {
			return memory[0]
		}
	}
	return memory[0]
}

func openFile() *os.File {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	return f
}

func parse(f *os.File) []int {
	scanner := bufio.NewScanner(f)
	tokens := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		tokens = append(tokens, strings.Split(text, ",")...)
	}
	numbers := []int{}
	for _, token := range tokens {
		numericToken, err := strconv.Atoi(token)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, numericToken)
	}
	return numbers
}
