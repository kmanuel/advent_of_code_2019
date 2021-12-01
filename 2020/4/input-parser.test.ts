import { readPassports } from './input-parser'

describe('readPassports', () => {
  it('should find passports', () => {
    const passports = readPassports('./4/sample.txt')
    expect(passports.length).toBe(4)
  })

  it('should parse passports', () => {
    const passports = readPassports('./4/sample.txt')
    expect(passports[0]).toEqual({
      fields: [
        'ecl:gry',
        'pid:860033327',
        'eyr:2020',
        'hcl:#fffffd',
        'byr:1937',
        'iyr:2017',
        'cid:147',
        'hgt:183cm',
      ],
    })
  })
})
