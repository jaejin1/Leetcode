package ValidAnagram

func isAnagram(s string, t string) bool {
	var xor, a, b rune

	for _, n := range s {
		xor ^= n
		a += n * n
	}

	for _, n := range t {
		xor ^= n
		b += n * n
	}

	return xor == 0 && a == b
}
