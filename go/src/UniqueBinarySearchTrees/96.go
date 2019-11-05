package UniqueBinarySearchTrees

func numTrees(n int) int {
	s := make([]int, n+1)
	s[0] = 1
	s[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			s[i] += s[j-1] * s[i-j]
		}
	}
	return s[n]
}
