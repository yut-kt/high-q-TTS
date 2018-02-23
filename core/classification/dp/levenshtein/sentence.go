package levenshtein

import (
	"strings"
	"sort"
)

type sentence struct {
	text string
}

func NewSentence(text string) *sentence {
	return &sentence{text: text}
}

func (s *sentence) Distance(text string) int {
	base, target := strings.Split(s.text, " "), strings.Split(text, " ")

	columnLen, rowLen := len(base)+1, len(target)+1

	matrix := make([][]int, columnLen)
	for column := 0; column < columnLen; column++ {
		matrix[column] = make([]int, len(target)+1)
	}

	for column := 1; column < columnLen; column++ {
		matrix[column][0] = matrix[column-1][0] + 1
	}
	for row := 1; row < rowLen; row++ {
		matrix[0][row] = matrix[0][row-1] + 1
	}

	for column := 1; column < columnLen; column++ {
		for row := 1; row < rowLen; row++ {
			cost := matrix[column-1][row-1] + map[bool]int{true: 0, false: 1}[base[column-1] == target[row-1]]
			costs := []int{cost, matrix[column-1][row] + 1, matrix[column][row-1] + 1}
			sort.Ints(costs)
			matrix[column][row] = costs[0]
		}
	}

	return matrix[columnLen-1][rowLen-1]
}
