import { countNumberIncreases, toThreeMeasureWindows } from './index'

describe('countNumberIncreases', () => {
  it('produces sample result', () => {
    const input = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263]
    const result = countNumberIncreases(input)
    expect(result).toBe(7)
  })

  it('maps measures to windows', () => {
    const input = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263]
    const result = toThreeMeasureWindows(input)
    expect(result).toEqual([607, 618, 618, 617, 647, 716, 769, 792])
  })
  it('detects measure sum increases', () => {
    const input = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263]
    const windows = toThreeMeasureWindows(input)
    const result = countNumberIncreases(windows)
    expect(result).toEqual(5)
  })
})
