package main

import (
	"errors"
)

type interpreter struct {
	memory             memory
	instructionPointer int
	fixedInputs        []int
	output             int
	inputCursor        int
	hasOutput          bool
	terminated         bool
}

func (i *interpreter) runForInput(input int) int {
	i.hasOutput = false
	i.fixedInputs = append(i.fixedInputs, input)
	i.calculateResult()
	return i.output
}

func (i *interpreter) addInput(input int) {
	i.fixedInputs = append(i.fixedInputs, input)
}

func (i *interpreter) isTerminateReached() bool {
	return i.instructionPointer >= len(i.memory) || i.terminated
}

func (i *interpreter) calculateResult() int {
	for !i.isTerminateReached() && !i.hasOutput {
		i.performNextInstruction()
	}
	return i.memory[0]
}

func (i *interpreter) performNextInstruction() {
	memory := i.memory
	instructionPointer := i.instructionPointer
	opCode := memory[instructionPointer]
	op, _ := i.getOperationForCode(opCode)
	res := op.execute(memory)
	i.instructionPointer = res
}

const (
	codeAdd         = 1
	codeMult        = 2
	codeInput       = 3
	codeOutput      = 4
	codeJumpIfTrue  = 5
	codeJumpIfFalse = 6
	codeLessThan    = 7
	codeEquals      = 8
	codeTerminate   = 99
)

func (i *interpreter) getOperationForCode(code int) (operation, error) {
	start := i.instructionPointer
	operationPart := code % 100
	if operationPart == codeAdd {
		return addOperation{start: start}, nil
	}
	if operationPart == codeMult {
		return multOperation{start: start}, nil
	}
	if operationPart == codeInput {
		input := i.fixedInputs[i.inputCursor]
		i.inputCursor = i.inputCursor + 1
		return &nonInteractiveInputOperation{
			start: start,
			input: input}, nil
	}
	if operationPart == codeOutput {
		return &outputOperation{start: start, output: &i.output, hasOutput: &i.hasOutput}, nil
	}
	if operationPart == codeTerminate {
		return terminateOperation{terminated: &i.terminated}, nil
	}
	if operationPart == codeJumpIfTrue {
		return jumpIfTrueOperation{start: start}, nil
	}
	if operationPart == codeJumpIfFalse {
		return jumpIfFalseOperation{start: start}, nil
	}
	if operationPart == codeLessThan {
		return lessThanOperation{start: start}, nil
	}
	if operationPart == codeEquals {
		return equalOperation{start: start}, nil
	}
	return nil, errors.New("no matching operation found")
}
