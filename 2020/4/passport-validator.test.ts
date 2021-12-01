import { readPassports } from './input-parser'
import { fieldValidations, isValid1, isValid2 } from './passport-validator'

const passports = readPassports('./4/sample.txt')

describe('isValid', () => {
  it('should detect valid passport', () => {
    expect(isValid1(passports[0])).toBeTruthy()
  })

  it('should detect invalid passport', () => {
    expect(isValid1(passports[1])).toBeFalsy()
  })

  it('should detect third sample as valid', () => {
    expect(isValid1(passports[2])).toBeTruthy()
  })

  it('should detect fourth sample as invalid', () => {
    expect(isValid1(passports[3])).toBeFalsy()
  })
})

const sample2Valid = readPassports('./4/sample2_valid.txt')
const sample2Invalid = readPassports('./4/sample2_invalid.txt')
describe('isValid2', () => {
  it('should detect valid passports', () => {
    expect(isValid2(sample2Valid[0])).toBeTruthy()
    expect(isValid2(sample2Valid[1])).toBeTruthy()
    expect(isValid2(sample2Valid[2])).toBeTruthy()
    expect(isValid2(sample2Valid[3])).toBeTruthy()
  })

  it('should detect invalid passports', () => {
    expect(isValid2(sample2Invalid[0])).toBeFalsy()
    expect(isValid2(sample2Invalid[1])).toBeFalsy()
    expect(isValid2(sample2Invalid[2])).toBeFalsy()
    expect(isValid2(sample2Invalid[3])).toBeFalsy()
  })
})

describe('validators', () => {
  it('ecl', () => {
    expect(fieldValidations.byr('2002')).toBeTruthy()
    expect(fieldValidations.byr('2003')).toBeFalsy()
  })
  it('byr', () => {
    expect(fieldValidations.hgt('60in')).toBeTruthy()
    expect(fieldValidations.hgt('190cm')).toBeTruthy()
    expect(fieldValidations.hgt('190in')).toBeFalsy()
    expect(fieldValidations.hgt('190')).toBeFalsy()
  })
  it('hcl', () => {
    expect(fieldValidations.hcl('#123abc')).toBeTruthy()
    expect(fieldValidations.hcl('#123abz')).toBeFalsy()
    expect(fieldValidations.hcl('123abc')).toBeFalsy()
  })
  it('ecl', () => {
    expect(fieldValidations.ecl('brn')).toBeTruthy()
    expect(fieldValidations.ecl('wat')).toBeFalsy()
  })
  it('pid', () => {
    expect(fieldValidations.pid('000000001')).toBeTruthy()
    expect(fieldValidations.pid('0123456789')).toBeFalsy()
  })
})
