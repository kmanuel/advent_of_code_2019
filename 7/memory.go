package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	refPosition  = 0
	refImmediate = 1
)

var memoryHolder memory

func newMemory() memory {
	if memoryHolder == nil {
		initializeMemoryHolder()
	}
	memoryCopy := make([]int, len(memoryHolder))
	copy(memoryCopy, memoryHolder)
	return memoryCopy
}

func initializeMemoryHolder() {
	f := openFile()
	defer f.Close()
	numbers := parse(f)
	memoryHolder = numbers
}

type memory []int

func (m memory) getReferencedValue(i int) int {
	return m[m[i]]
}

func (m memory) getImmediateValue(i int) int {
	return m[i]
}

func (m memory) getReference(i int) *int {
	return &m[m[i]]
}

func (m memory) getArgumentNr(opStart int, argNum int) int {
	op := m[opStart]
	var refType int
	offset := 10
	for i := 0; i < argNum; i++ {
		offset *= 10
	}
	refType = op / offset % 10
	if refType == refPosition {
		return m.getReferencedValue(opStart + argNum)
	}
	return m.getImmediateValue(opStart + argNum)
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
