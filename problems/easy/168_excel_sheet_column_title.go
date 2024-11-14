package easy

import (
	"fmt"
	"leetcode/solutions/easy"
)

/**
Given an integer columnNumber, return its corresponding column title as it appears in an Excel sheet.

For example:

A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...


Example 1:

Input: columnNumber = 1
Output: "A"
*/

func GetColumnTitle() {
	columnNumber := 28

	column := easy.ConvertToTitle(columnNumber)

	fmt.Println(column)
}
