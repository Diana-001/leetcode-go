package easy

import "strings"

func SolutionOfProblem58(s string) int {
	arr := strings.Fields(s)
	return len((arr[len(arr)-1]))
}
