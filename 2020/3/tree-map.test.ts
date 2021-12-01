import { TreeMap } from './tree-map'

describe('TreeMap', () => {
  it('should parse file without error', () => {
    const treeMap = TreeMap.fromFile('./3/sample.txt')
    expect(TreeMap).toBeDefined()
  })

  it('should set tree on position 0,2', () => {
    const treeMap = TreeMap.fromFile('./3/sample.txt')
    expect(treeMap.isTree(2, 0)).toBeTruthy()
  })

  it('should not set tree on position 0,0', () => {
    const treeMap = TreeMap.fromFile('./3/sample.txt')
    expect(treeMap.isTree(0, 0)).toBeFalsy()
  })

  it('should use repeated pattern if out of bounds of input', () => {
    const treeMap = TreeMap.fromFile('./3/sample.txt')
    expect(treeMap.isTree(11, 0)).toBeFalsy()
    expect(treeMap.isTree(12, 0)).toBeFalsy()
    expect(treeMap.isTree(13, 0)).toBeTruthy()
  })
})
