package TargetSum

func findTargetSumWays(nums []int, S int) int {
	var s []int
	var temp int
	var result int
	var dp map[int][]int

	dp = map[int][]int{}

	for i, n := range nums {
		if i == 0 {
			s = append(s, n)
		} else {
			temp = len(s)
			for j, m := range s {
				if _, ok := dp[m]; ok {
					s = append(s, dp[m]...)
					if temp-1 == j {
						s = s[temp:]
					}
				} else {
					s = append(s, m+n)
					s = append(s, m-n)
					if temp-1 == j {
						s = s[temp:]
					}
					dpslice := []int{m + n, m - n}
					dp[m] = append(dp[m], dpslice[0])
					dp[m] = append(dp[m], dpslice[1])
				}

			}
			for k := range dp {
				delete(dp, k)
			}
		}

	}

	for _, item := range s {
		if item == S || item == -1*S {
			if item == 0 {
				result += 2
			} else {
				result++
			}
		}
	}

	return result
}
