import fs = require('fs')
import {
  countValidPasswords1,
  countValidPasswords2,
} from './password-policy-checker'

export const loadInput = () => {
  const data = fs.readFileSync('./2/input.txt', 'utf-8')
  return data.split('\n')
}

const run1 = () => {
  const input = loadInput()
  return countValidPasswords1(input)
}

const run2 = () => {
  const input = loadInput()
  return countValidPasswords2(input)
}

const output = run1()
console.log(`Day 2, challenge 1: ${output}`)
console.log(`Day 2, challenge 2: ${run2()}`)
