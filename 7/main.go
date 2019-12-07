package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	res := getMaxResultForDifferentPhaseSettings()
	fmt.Printf("%d\n", res)
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}

func getMaxResultForDifferentPhaseSettings() int {
	max := 0
	phaseSettings := createPhaseSettings(5)
	for _, phaseSetting := range phaseSettings {
		res := runPhaseSetting(phaseSetting)
		if res > max {
			fmt.Printf("%v produced a new max of %d\n", phaseSetting, res)
			max = res
		}
	}
	return max
}

func runPhaseSetting(phaseSetting []int) int {
	a := newInterpreter(phaseSetting[0])
	b := newInterpreter(phaseSetting[1])
	c := newInterpreter(phaseSetting[2])
	d := newInterpreter(phaseSetting[3])
	e := newInterpreter(phaseSetting[4])
	res := 0
	running := true
	for running {
		res = a.runForInput(res)
		if a.terminated {
			break
		}
		res = b.runForInput(res)
		if b.terminated {
			break
		}
		res = c.runForInput(res)
		if c.terminated {
			break
		}
		res = d.runForInput(res)
		if d.terminated {
			break
		}
		res = e.runForInput(res)
		if e.terminated {
			break
		}
	}
	return e.output
}

func newInterpreter(phaseSetting int) interpreter {
	memory := newMemory()
	interpreter := interpreter{
		memory:      memory,
		fixedInputs: []int{phaseSetting},
	}
	return interpreter
}

func star2() int {

	return -1
}
