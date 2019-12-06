package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	id   string
	prev *node
}

func (n node) getDistance() int {
	if n.prev == nil {
		return 0
	}
	return n.prev.getDistance() + 1
}

func (n node) getRootPath() []*node {
	path := []*node{}
	curr := &n
	for curr != nil {
		path = append(path, curr)
		curr = curr.prev
	}
	return path
}

func main() {
	star2()
}

func star1() {
	f := openFile()
	defer f.Close()
	tokens := parse(f)
	nodes := buildNodeMap(tokens)
	sum := 0
	for _, v := range nodes {
		sum += v.getDistance()
	}
	fmt.Printf("total orbits %d\n", sum)
}

func star2() {
	f := openFile()
	defer f.Close()
	tokens := parse(f)
	nodes := buildNodeMap(tokens)
	steps := getStepsBetweenYouAndSanta(nodes)
	print(steps - 2)
}

func getStepsBetweenYouAndSanta(nodes map[string]*node) int {
	santaPath := nodes["SAN"].getRootPath()
	youPath := nodes["YOU"].getRootPath()
	for youSteps, yNode := range youPath {
		for santaSteps, sNode := range santaPath {
			if yNode == sNode {
				return youSteps + santaSteps
			}
		}
	}
	return -1
}

func buildNodeMap(tokens []string) map[string]*node {
	nodes := map[string]*node{}

	for _, token := range tokens {
		split := strings.Split(token, ")")
		first := split[0]
		second := split[1]
		n1 := createNode(nil, first)
		n2 := createNode(nil, second)
		nodes[first] = &n1
		nodes[second] = &n2
	}
	for _, token := range tokens {
		split := strings.Split(token, ")")
		first := split[0]
		second := split[1]
		prevNode := nodes[first]
		currNode := nodes[second]
		currNode.prev = prevNode
	}
	return nodes
}

func createNode(prevNode *node, id string) node {
	if prevNode != nil {
		return node{
			id:   id,
			prev: prevNode,
		}
	}
	return node{
		id:   id,
		prev: nil,
	}
}

func openFile() *os.File {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	return f
}

func parse(f *os.File) []string {
	scanner := bufio.NewScanner(f)
	tokens := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		tokens = append(tokens, strings.Split(text, "\n")...)
	}
	return tokens
}
