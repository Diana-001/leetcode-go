package tests

import (
	"leetcode/solutions/easy"
	"testing"
)

// Тест для функции ConvertToTitle
func TestConvertToTitle(t *testing.T) {
	tests := []struct {
		columnNumber int
		expected     string
	}{
		{1, "A"},
		{26, "Z"},
		{27, "AA"},
		{52, "AZ"},
		{701, "ZY"},
		{702, "ZZ"},
		{703, "AAA"},
	}

	for _, test := range tests {
		t.Run("Test", func(t *testing.T) {
			result := easy.ConvertToTitle(test.columnNumber)

			if result != test.expected {
				t.Errorf("For input %d, expected %s, but got %s", test.columnNumber, test.expected, result)
			}
		})
	}
}
