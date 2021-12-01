import fs = require('fs')

export const loadInput = (fileName: string) => {
  const data = fs.readFileSync(fileName, 'utf-8')
  return data.split('\n').map((r) => Number(r))
}
