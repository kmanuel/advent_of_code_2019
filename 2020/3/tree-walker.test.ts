import { walkTree } from './tree-walker'
describe('walkTree', () => {
  it('should return 7 on sample input', () => {
    const trees = walkTree('./3/sample.txt', 3, 1)
    expect(trees).toBe(7)
  })
})
