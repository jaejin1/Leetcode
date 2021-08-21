package RemoveOutermostParentheses

import (
	"bytes"
	"fmt"
)

func removeOuterParentheses(S string) string {
	var res bytes.Buffer
	counter := 0
	for _, c := range S {
		fmt.Println(c, counter, res.String())
		if counter != 0 && !(counter == 1 && c == ')') {
			res.WriteRune(c)
		}
		if c == '(' {
			counter++
		} else {
			counter--
		}

	}
	return res.String()
}
