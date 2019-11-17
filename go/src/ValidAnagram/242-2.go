package ValidAnagram

func isAnagram(s string, t string) bool {
	var a map[rune]int
	a = map[rune]int{}

	for _, n := range s {
		a[n]++
	}

	for _, n := range t {
		if _, ok := a[n]; ok {
			a[n]--
		} else {
			return false
		}
	}

	for _, n := range a {
		if n != 0 {
			return false
		}
	}

	return true
}
