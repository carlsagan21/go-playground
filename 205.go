package main

func isIsomorphic(s string, t string) bool {
	dic := make(map[rune]rune)
	revDic := make(map[rune]rune)
	for i, sc := range []rune(s) {
		tc := rune(t[i])
		if tcInDic, ok := dic[sc]; ok {
			if tcInDic != tc {
				return false
			}
		} else {
			if _, ok := revDic[tc]; ok {
				return false
			}
			revDic[tc] = sc
			dic[sc] = tc
		}
	}
	return true
}
