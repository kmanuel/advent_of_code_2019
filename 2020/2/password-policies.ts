import {
  PasswordCheckerFunction,
  PasswordCheckerOptions,
} from './password-policy-checker'

export const occurrenceCheck: PasswordCheckerFunction = (
  { fromNum, toNum, letter }: PasswordCheckerOptions,
  password: string
): boolean => {
  const occurrences = password
    .split('')
    .filter((l: string) => l === letter).length
  return occurrences >= fromNum && occurrences <= toNum
}

export const positionCheck: PasswordCheckerFunction = (
  { fromNum, toNum, letter },
  password
) => {
  const first = password[fromNum - 1] === letter
  const second = password[toNum - 1] === letter
  return (first && !second) || (!first && second)
}
