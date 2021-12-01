import { loadInput } from './input-parser'

describe('input-parser', () => {
  it('reads file', () => {
    const result = loadInput('./1/test.txt')
    expect(result).toEqual([1, 2, 3])
  })
})
