package main

import (
	"fmt"
)

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
	terminated *bool
}

func (o terminateOperation) execute(memory memory) int {
	*o.terminated = true
	return -1
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

type nonInteractiveInputOperation struct {
	start int
	input int
}

func (o *nonInteractiveInputOperation) execute(memory memory) int {
	s := o.start
	targetRef := memory.getReference(s + 1)
	*targetRef = o.input
	return s + 2
}

type outputOperation struct {
	start     int
	output    *int
	hasOutput *bool
}

func (o *outputOperation) execute(m memory) int {
	s := o.start
	value := m.getArgumentNr(s, 1)
	*o.output = value
	*o.hasOutput = true
	return s + 2
}
