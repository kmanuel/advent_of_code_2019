package main

func createPhaseSettings(start int) [][]int {
	phaseSettings := [][]int{}
	for f1 := 0; f1 < 5; f1++ {
		for f2 := 0; f2 < 5; f2++ {
			if f1 == f2 {
				continue
			}
			for f3 := 0; f3 < 5; f3++ {
				if f3 == f1 || f3 == f2 {
					continue
				}
				for f4 := 0; f4 < 5; f4++ {
					if f4 == f3 || f4 == f2 || f4 == f1 {
						continue
					}
					for f5 := 0; f5 < 5; f5++ {
						if f5 == f4 || f5 == f3 || f5 == f2 || f5 == f1 {
							continue
						}
						phaseSettings = append(phaseSettings, []int{start + f1, start + f2, start + f3, start + f4, start + f5})
					}
				}
			}
		}
	}
	return phaseSettings
}
