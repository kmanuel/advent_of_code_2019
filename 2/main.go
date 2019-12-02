package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func main() {
	println(star1())
	println(star2())
}

func star1() int {
	f := openFile()
	defer f.Close()
	numbers := parse(f)
	result := calculateResult(numbers)
	return result
}

func star2() int {
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
				return endRes
			}
			copy(inputCopy, numbers)
		}
	}
	return -1
}

const (
	OPCODE_ADD       = 1
	OPCODE_MULT      = 2
	OPCODE_TERMINATE = 99
)

func calculateResult(memory []int) int {
	for i := 0; i < len(memory); {
		opCode := memory[i]
		op, err := getOperationForCode(opCode, i)
		if err != nil {
			return -1
		}
		res := op.execute(memory)
		if res == 0 {
			return memory[0]
		}
		i += res
	}
	return memory[0]
}

func getOperationForCode(code int, start int) (operation, error) {
	if code == OPCODE_ADD {
		return addOperation{start: start}, nil
	}
	if code == OPCODE_MULT {
		return multOperation{start: start}, nil
	}
	if code == OPCODE_TERMINATE {
		return terminateOperation{}, nil
	}
	return nil, errors.New("no matching operation found")
}

type memory []int

func (m memory) getReferencedValue(i int) int {
	return m[m[i]]
}

func (m memory) getReference(i int) *int {
	return &m[m[i]]
}

type operation interface {
	execute(memory memory) int
}

type addOperation struct {
	start int
}

func (o addOperation) execute(m memory) int {
	s := o.start
	val1 := m.getReferencedValue(s + 1)
	val2 := m.getReferencedValue(s + 2)
	targetRef := m.getReference(s + 3)
	*targetRef = val1 + val2
	return 4
}

type multOperation struct {
	start int
}

func (o multOperation) execute(m memory) int {
	s := o.start
	val1 := m.getReferencedValue(s + 1)
	val2 := m.getReferencedValue(s + 2)
	targetRef := m.getReference(s + 3)
	*targetRef = val1 * val2
	return 4
}

type terminateOperation struct {
}

func (o terminateOperation) execute(memory memory) int {
	return 0
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
