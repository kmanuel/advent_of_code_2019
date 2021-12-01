import * as passwordPolicies from './password-policies'

export type PasswordCheckerFunction = (
  options: PasswordCheckerOptions,
  password: string
) => boolean

export interface PasswordCheckerOptions {
  fromNum: number
  toNum: number
  letter: string
}

const parseLine = (input: string) => {
  const match = input.match(/^(\d+)-(\d+) (\S): (.*)$/)
  if (!match) {
    throw new Error('')
  }
  const fromNum = Number(match[1])
  const toNum = Number(match[2])
  const letter = match[3]
  const password = match[4]
  return {
    fromNum,
    toNum,
    letter,
    password,
  }
}

export const isValidPassword = (
  input: string,
  checker: PasswordCheckerFunction
): boolean => {
  const parsed = parseLine(input)
  return checker(parsed, parsed.password)
}

export const countValidPasswords1 = (inputLines: string[]): number => {
  return inputLines.filter((i) =>
    isValidPassword(i, passwordPolicies.occurrenceCheck)
  ).length
}

export const countValidPasswords2 = (inputLines: string[]): number => {
  return inputLines.filter((i) =>
    isValidPassword(i, passwordPolicies.positionCheck)
  ).length
}
