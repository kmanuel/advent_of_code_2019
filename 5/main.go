package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	star1()
	// println(star2())
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
	OPCODE_ADD          = 1
	OPCODE_MULT         = 2
	OPCODE_INPUT        = 3
	OPCODE_OUTPUT       = 4
	OPCODE_JUMP_IF_TRUE = 5
	OPCODE_JUMP_IF_FALSE = 6
	OPCODE_LESS_THAN = 7
	OPCODE_EQUALS = 8
	OPCODE_TERMINATE    = 99
	REFTYPE_POSITION  = 0
	REFTYPE_IMMEDIATE = 1
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
		i = res
	}
	return memory[0]
}

func getOperationForCode(code int, start int) (operation, error) {
	operationPart := code % 100
	if operationPart == OPCODE_ADD {
		return addOperation{start: start}, nil
	}
	if operationPart == OPCODE_MULT {
		return multOperation{start: start}, nil
	}
	if operationPart == OPCODE_INPUT {
		return inputOperation{start: start}, nil
	}
	if operationPart == OPCODE_OUTPUT {
		return outputOperation{start: start}, nil
	}
	if operationPart == OPCODE_TERMINATE {
		return terminateOperation{}, nil
	}
	if operationPart == OPCODE_JUMP_IF_TRUE {
		return jumpIfTrueOperation{start: start}, nil
	}
	if operationPart == OPCODE_JUMP_IF_FALSE {
		return jumpIfFalseOperation{start: start}, nil
	}
	if operationPart == OPCODE_LESS_THAN {
		return lessThanOperation{start: start}, nil
	}
	if operationPart == OPCODE_EQUALS {
		return equalOperation{start: start}, nil
	}
	return nil, errors.New("no matching operation found")
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
	if argNum == 1 {
		refType = op / 100 % 10
	}
	if argNum == 2 {
		refType = op / 1000 % 10
	}
	if argNum == 3 {
		refType = op / 10000 % 10
	}
	if refType == REFTYPE_POSITION {
		return m.getReferencedValue(opStart + argNum)
	}
	return m.getImmediateValue(opStart + argNum)
}

type operation interface {
	execute(memory memory) int
}

type addOperation struct {
	start int
}

func (o addOperation) execute(m memory) int {
	s := o.start
	val1 := m.getArgumentNr(s, 1)
	val2 := m.getArgumentNr(s, 2)
	targetRef := m.getReference(s + 3)
	*targetRef = val1 + val2
	return s + 4
}

type multOperation struct {
	start int
}

func (o multOperation) execute(m memory) int {
	s := o.start
	val1 := m.getArgumentNr(s, 1)
	val2 := m.getArgumentNr(s, 2)
	targetRef := m.getReference(s + 3)
	*targetRef = val1 * val2
	return s + 4
}

type jumpIfTrueOperation struct {
	start int
}

func (o jumpIfTrueOperation) execute(m memory) int {
	s := o.start
	val1 := m.getArgumentNr(s, 1)
	val2 := m.getArgumentNr(s, 2)
	if val1 != 0 {
		return val2
	}
	return s + 3
}

type jumpIfFalseOperation struct {
	start int
}

func (o jumpIfFalseOperation) execute(m memory) int {
	s := o.start
	val1 := m.getArgumentNr(s, 1)
	val2 := m.getArgumentNr(s, 2)
	if val1 == 0 {
		return val2
	}
	return s + 3
}

type lessThanOperation struct {
	start int
}

func (o lessThanOperation) execute(m memory) int {
	s := o.start
	val1 := m.getArgumentNr(s, 1)
	val2 := m.getArgumentNr(s, 2)
	if val1 < val2 {
		targetRef := m.getReference(s + 3)
		*targetRef = 1
	} else {
		targetRef := m.getReference(s + 3)
		*targetRef = 0
	}
	return s + 4
}

type equalOperation struct {
	start int
}

func (o equalOperation) execute(m memory) int {
	s := o.start
	val1 := m.getArgumentNr(s, 1)
	val2 := m.getArgumentNr(s, 2)
	if val1 == val2 {
		targetRef := m.getReference(s + 3)
		*targetRef = 1
	} else {
		targetRef := m.getReference(s + 3)
		*targetRef = 0
	}
	return s + 4
}

type terminateOperation struct {
}

func (o terminateOperation) execute(memory memory) int {
	return 0
}

type inputOperation struct {
	start int
}

func (o inputOperation) execute(memory memory) int {
	fmt.Print("Enter: ")
	var input int
	fmt.Scan(&input)
	s := o.start
	targetRef := memory.getReference(s + 1)
	*targetRef = input
	return s + 2
}

type outputOperation struct {
	start int
}

func (o outputOperation) execute(m memory) int {
	s := o.start
	value := m.getArgumentNr(s, 1)
	println(value)
	return s + 2
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
