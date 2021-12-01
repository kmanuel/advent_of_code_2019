import { readPassports } from './input-parser'
import { isValid1, isValid2 } from './passport-validator'
const passports = readPassports('./4/input.txt')

const validPassportsCount = passports.filter(isValid1).length
console.log(`Day 4, Part 1: ${validPassportsCount}`)
console.log(`Day 4, Part 2: ${passports.filter(isValid2).length}`)
