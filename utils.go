package main

func mergeSkillMaps(m map[string]int, ms ...map[string]int) map[string]int {
	for _, mm := range ms {
		for skill, level := range mm {
			if _, ok := m[skill]; ok {
				m[skill] += level
			} else {
				m[skill] = level
			}
		}
	}
	return m
}
