import { findNNumbersSummingUpTo, multiplyAll, getNMult } from './solver'

describe('solver', () => {
  it('findNNumbersSummingUpTo one number', () => {
    const result = findNNumbersSummingUpTo(1, [1, 2, 3, 4, 5, 6], 5)
    expect(result).toEqual([5])
  })
  it('findSumOfTwoNumbersAddingUpTo should return correct number for input', () => {
    const result = findNNumbersSummingUpTo(
      2,
      [1721, 979, 366, 299, 675, 1456],
      2020
    )
    expect(result).toEqual([1721, 299])
  })

  it('multiplyAll should multiply', () => {
    expect(multiplyAll([1721, 299])).toBe(514579)
  })

  it('getNMult 2 nums', () => {
    expect(getNMult(2, [1721, 979, 366, 299, 675, 1456])).toBe(514579)
  })

  it('getNMult 3 nums', () => {
    expect(getNMult(3, [1721, 979, 366, 299, 675, 1456])).toBe(241861950)
  })
})
