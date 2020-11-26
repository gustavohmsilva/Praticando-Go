package namelist

import "sort"

func increaseMap(m map[string]int, s []string) map[string]int {
	for _, v := range s {
		_, exist := m[v]
		if !exist {
			m[v] = 1
		} else {
			m[v]++
		}
	}
	return m
}

func merge(s1, s2 []string) []string {
	uniqueResults := countRepeated(s1, s2)
	resultingSlice := newStringSlice(uniqueResults)
	sort.Strings(resultingSlice)
	return resultingSlice
}

func newStringSlice(m map[string]int) []string {
	var r []string
	for k := range m {
		r = append(r, k)
	}
	return r
}

func countRepeated(s1, s2 []string) map[string]int {
	m := make(map[string]int)
	m = increaseMap(m, s1)
	m = increaseMap(m, s2)
	return m
}
