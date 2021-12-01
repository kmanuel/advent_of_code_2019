import fs from 'fs'
const filePath = './1/input.txt'
function parseInput(filePath: string) {
  return fs
    .readFileSync(filePath, 'utf-8')
    .split('\n')
    .map((n) => Number(n))
}

const solve1 = () => {
  const nums = parseInput(filePath)
  return countNumberIncreases(nums)
}

const solve2 = () => {
  const nums = parseInput(filePath)
  const windows = toThreeMeasureWindows(nums)
  return countNumberIncreases(windows)
}

export const countNumberIncreases = (nums: number[]): number => {
  let count = 0
  for (let i = 1; i < nums.length; i++) {
    if (nums[i - 1] < nums[i]) {
      count++
    }
  }
  return count
}

export const toThreeMeasureWindows = (nums: number[]): number[] => {
  const measures = nums
    .slice(0, nums.length - 2)
    .map((n, idx) => n + nums[idx + 1] + nums[idx + 2])
  return measures
}

console.log(`Day1 p1= ${solve1()}`)
console.log(`Day1 p2= ${solve2()}`)
