import fs from 'fs'

export interface Passport {
  fields: string[]
}

const passportFrom = (passportLines: string) => {
  const fields = passportLines.split(/\s/).filter((s) => s.trim().length > 0)
  return { fields }
}

export const readPassports = (fileName: string): Passport[] => {
  const fileContent = fs.readFileSync(fileName, 'utf-8')
  const passportLines = fileContent.split(/^\n/gm)
  return passportLines.map((lines) => passportFrom(lines))
}
