package main

// 771. Jewels and Stones
// Easy
// https://leetcode.com/problems/jewels-and-stones/
func numJewelsInStones(J string, S string) int {
	jewels := asRuneMap(J)
	return countRunesInMap(S, jewels)
}

func countRunesInMap(S string, m map[rune]struct{}) int {
	counter := 0
	for _, s := range []rune(S) {
		if _, ok := m[s]; ok {
			counter++
		}
	}
	return counter
}

func asRuneMap(J string) map[rune]struct{} {
	var jewels = make(map[rune]struct{})
	for _, j := range []rune(J) {
		jewels[j] = struct{}{}
	}
	return jewels
}
