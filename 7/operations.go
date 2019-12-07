package main

import (
	"fmt"
)

type operation interface {
	execute(memory memory, state *state)
}

type addOperation struct{}

func (o addOperation) execute(m memory, state *state) {
	s := state.instructionPointer
	val1 := m.getArgumentNr(s, 1)
	val2 := m.getArgumentNr(s, 2)
	targetRef := m.getReference(s + 3)
	*targetRef = val1 + val2
	state.moveInstructionPointer(4)
}

type multOperation struct{}

func (o multOperation) execute(m memory, state *state) {
	s := state.instructionPointer
	val1 := m.getArgumentNr(s, 1)
	val2 := m.getArgumentNr(s, 2)
	targetRef := m.getReference(s + 3)
	*targetRef = val1 * val2
	state.moveInstructionPointer(4)
}

type jumpIfTrueOperation struct {
}

func (o jumpIfTrueOperation) execute(m memory, state *state) {
	s := state.instructionPointer
	val1 := m.getArgumentNr(s, 1)
	val2 := m.getArgumentNr(s, 2)
	if val1 != 0 {
		state.setInstructionPointer(val2)
		return
	}
	state.moveInstructionPointer(3)
}

type jumpIfFalseOperation struct{}

func (o jumpIfFalseOperation) execute(m memory, state *state) {
	s := state.instructionPointer
	val1 := m.getArgumentNr(s, 1)
	val2 := m.getArgumentNr(s, 2)
	if val1 == 0 {
		state.setInstructionPointer(val2)
		return
	}
	state.moveInstructionPointer(3)
}

type lessThanOperation struct{}

func (o lessThanOperation) execute(m memory, state *state) {
	s := state.instructionPointer
	val1 := m.getArgumentNr(s, 1)
	val2 := m.getArgumentNr(s, 2)
	if val1 < val2 {
		targetRef := m.getReference(s + 3)
		*targetRef = 1
	} else {
		targetRef := m.getReference(s + 3)
		*targetRef = 0
	}
	state.moveInstructionPointer(4)
}

type equalOperation struct{}

func (o equalOperation) execute(m memory, state *state) {
	s := state.instructionPointer
	val1 := m.getArgumentNr(s, 1)
	val2 := m.getArgumentNr(s, 2)
	if val1 == val2 {
		targetRef := m.getReference(s + 3)
		*targetRef = 1
	} else {
		targetRef := m.getReference(s + 3)
		*targetRef = 0
	}
	state.moveInstructionPointer(4)
}

type terminateOperation struct{}

func (o terminateOperation) execute(memory memory, state *state) {
	state.terminated = true
	state.setInstructionPointer(-1)
}

type inputOperation struct{}

func (o inputOperation) execute(memory memory, state *state) {
	fmt.Print("Enter: ")
	var input int
	fmt.Scan(&input)
	s := state.instructionPointer
	targetRef := memory.getReference(s + 1)
	*targetRef = input
	state.moveInstructionPointer(2)
}

type nonInteractiveInputOperation struct {
}

func (o *nonInteractiveInputOperation) execute(memory memory, state *state) {
	s := state.instructionPointer
	targetRef := memory.getReference(s + 1)
	*targetRef = state.getNextInput()
	state.moveInstructionPointer(2)
}

type outputOperation struct{}

func (o *outputOperation) execute(m memory, state *state) {
	s := state.instructionPointer
	value := m.getArgumentNr(s, 1)
	state.addOutput(value)
	state.moveInstructionPointer(2)
}
