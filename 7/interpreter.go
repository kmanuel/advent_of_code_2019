package main

type interpreter struct {
	memory memory
	state  state
}

func newInterpreter(phaseSetting int) interpreter {
	memory := newMemory()
	interpreter := interpreter{
		memory: memory,
		state:  state{},
	}
	interpreter.setPhase(phaseSetting)
	return interpreter
}

func (i *interpreter) runForInput(input int) int {
	i.state.hasOutput = false
	i.state.fixedInputs = append(i.state.fixedInputs, input)
	i.calculateResult()
	return i.state.output
}

func (i *interpreter) setPhase(input int) {
	i.state.addInput(input)
}

func (i *interpreter) isTerminateReached() bool {
	return i.state.instructionPointer >= len(i.memory) || i.state.terminated
}

func (i *interpreter) calculateResult() int {
	for !i.isTerminateReached() && !i.state.hasOutput {
		op := i.getNextOperation()
		op.execute(i.memory, &i.state)
	}
	return i.memory[0]
}

func (i *interpreter) getNextOperation() operation {
	instructionPointer := i.state.instructionPointer
	opCode := i.memory[instructionPointer]
	operationPart := opCode % 100
	return codesToOperations[operationPart]
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

var codesToOperations = map[int]operation{
	codeAdd:         addOperation{},
	codeMult:        multOperation{},
	codeInput:       &nonInteractiveInputOperation{},
	codeOutput:      &outputOperation{},
	codeJumpIfTrue:  &jumpIfTrueOperation{},
	codeJumpIfFalse: jumpIfFalseOperation{},
	codeLessThan:    lessThanOperation{},
	codeEquals:      equalOperation{},
	codeTerminate:   terminateOperation{},
}
