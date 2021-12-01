import { Passport } from './input-parser'

const requiredFields = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid']

export const isValid1 = (passport: Passport): boolean => {
  for (let field of requiredFields) {
    if (!passport.fields.map((f) => f.split(':')[0]).includes(field)) {
      return false
    }
  }
  return true
}

const numberBetween = (value: string, from: number, to: number) => {
  const val = Number(value)
  return val >= from && val <= to
}

export const fieldValidations = {
  byr: (value: string) => numberBetween(value, 1920, 2002),
  iyr: (value: string) => numberBetween(value, 2010, 2020),
  eyr: (value: string) => numberBetween(value, 2020, 2030),
  hgt: (value: string) => {
    const match = value.match(/(\d+)(.*)/)
    const val = Number(match![1])
    const type = match![2]
    if (type === 'cm') {
      return val >= 150 && val <= 193
    } else if (type === 'in') {
      return val >= 59 && val <= 76
    }
    return false
  },
  hcl: (value: string) => {
    return value.match(/^#[a-f0-9]{6}$/)
  },
  ecl: (value: string) => {
    return value.match(/^(amb|blu|brn|gry|grn|hzl|oth)$/)
  },
  pid: (value: string) => {
    return value.match(/^[0-9]{9}$/)
  },
}

export const isValid2 = (passport: Passport): boolean => {
  if (!isValid1(passport)) {
    return false
  }
  for (let field of passport.fields) {
    const parts = field.split(':')
    if (Object.keys(fieldValidations).includes(parts[0])) {
      const key = parts[0] as
        | 'byr'
        | 'iyr'
        | 'eyr'
        | 'hgt'
        | 'hcl'
        | 'ecl'
        | 'pid'
      const validField = fieldValidations[key](parts[1])
      if (!validField) {
        return false
      }
    }
  }
  return true
}
