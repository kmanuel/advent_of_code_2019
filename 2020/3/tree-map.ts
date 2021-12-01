import fs from 'fs'

const readInput = (filePath: string) => {
  return fs.readFileSync(filePath, 'utf-8').split('\n')
}

export class TreeMap {
  inputLines: string[]

  isTree(row: number, col: number): boolean {
    const column = this.inputLines[col]
    const colSplit = column.split('')
    const foundSymbol = colSplit[row % colSplit.length]
    return foundSymbol === '#'
  }

  getTreeHeight(): number {
    return this.inputLines.length
  }

  constructor(inputLines: string[]) {
    this.inputLines = inputLines
  }

  static fromFile(filePath: string): TreeMap {
    const input = readInput(filePath)
    return new TreeMap(input)
  }
}
