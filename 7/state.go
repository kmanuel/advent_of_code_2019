package main

type state struct {
	instructionPointer int
	fixedInputs        []int
	output             int
	inputCursor        int
	hasOutput          bool
	terminated         bool
}

func (s *state) addInput(n int) {
	s.fixedInputs = append(s.fixedInputs, n)
}

func (s *state) getNextInput() int {
	val := s.fixedInputs[s.inputCursor]
	s.inputCursor++
	return val
}

func (s *state) addOutput(n int) {
	s.output = n
	s.hasOutput = true
}

func (s *state) moveInstructionPointer(n int) {
	s.instructionPointer += n
}

func (s *state) setInstructionPointer(n int) {
	s.instructionPointer = n
}
