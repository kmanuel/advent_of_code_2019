import { loadInput } from './input-parser'
import { getNMult } from './solver'

const input = loadInput('./1/input.txt')
const p1Result = getNMult(2, input)
console.log(`Part One result: ${p1Result}`)
const p2Result = getNMult(3, input)
console.log(`Part Two result: ${p2Result}`)
