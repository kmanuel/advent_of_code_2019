import { TreeMap } from './tree-map'

export const walkTree = (
  inputFile: string,
  rowMove: number,
  colMove: number
) => {
  let colPos = 0
  let rowPos = 0
  const tree = TreeMap.fromFile(inputFile)
  let treesEncountered = 0
  while (colPos < tree.getTreeHeight()) {
    if (tree.isTree(rowPos, colPos)) {
      treesEncountered++
    }
    colPos += colMove
    rowPos += rowMove
  }
  return treesEncountered
}
