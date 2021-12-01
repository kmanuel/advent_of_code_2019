import { walkTree } from './tree-walker'

const trees = walkTree('./3/sample.txt', 3, 1)
console.log(`Sample result = ${trees}`)

const ex1Trees = walkTree('./3/input.txt', 3, 1)
console.log(`Day 3, Part 1: ${ex1Trees}`)

const getMultiSlopeRes = () => {
  return (
    walkTree('./3/input.txt', 1, 1) *
    walkTree('./3/input.txt', 3, 1) *
    walkTree('./3/input.txt', 5, 1) *
    walkTree('./3/input.txt', 7, 1) *
    walkTree('./3/input.txt', 1, 2)
  )
}

console.log(`Day 3, Part 2: ${getMultiSlopeRes()}`)
