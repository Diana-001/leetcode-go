package tests

import (
	"strings"
	"testing"
)

func TestProblem58(t *testing.T) {
	s := "luffy is still joyboy  "
	arr := strings.Fields(s)

	findLengthOfLastWord := len((arr[len(arr)-1]))

	lengthOfLastWord := 6

	if findLengthOfLastWord != lengthOfLastWord {
		t.Errorf("got %q, wanted %q", arr, lengthOfLastWord)
	}
}
