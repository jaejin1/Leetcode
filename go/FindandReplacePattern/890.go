package FindandReplacePattern

func findAndReplacePattern(words []string, pattern string) []string {
	var result []string
	var check bool = true

	for _, value := range words {
		m1 := make(map[string]string)
		m2 := make(map[string]string)
		for j, pat := range pattern {

			if _, ok := m1[string(value[j])]; !ok {
				m1[string(value[j])] = string(pat)
			}

			if _, ok := m2[string(pat)]; !ok {
				m2[string(pat)] = string(value[j])
			}

			if m1[string(value[j])] != string(pat) {
				check = false
			}

			if m2[string(pat)] != string(value[j]) {
				check = false
			}
		}
		if check {
			result = append(result, value)
		}
		check = true
	}
	return result
}
