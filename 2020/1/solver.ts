export const findNNumbersSummingUpTo = (
  n: number,
  numbers: number[],
  to: number
): number[] => {
  if (n <= 0) {
    return []
  }
  if (numbers.length === 0) {
    return []
  }
  for (let i = 0; i < numbers.length - (n - 1); i++) {
    const currSum = numbers[i]
    if (n === 1 && currSum === to) {
      return [numbers[i]]
    }
    if (n - 1 > 0 && currSum < to) {
      const foundNumbers = findNNumbersSummingUpTo(
        n - 1,
        numbers.slice(i + 1),
        to - currSum
      )
      if (foundNumbers.length === n - 1) {
        return [numbers[i], ...foundNumbers]
      }
    }
  }
  return []
}

export const multiplyAll = (nums: number[]): number =>
  nums.reduce((acc, num) => acc * num, 1)

export const getNMult = (n: number, numbers: number[]) => {
  const nums = findNNumbersSummingUpTo(n, numbers, 2020)
  return multiplyAll(nums)
}
