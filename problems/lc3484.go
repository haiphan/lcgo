package problems

import (
	"strconv"
	"strings"
	"unicode"
)

type Spreadsheet struct {
	m map[string]int
}

func SpreadsheetConstructor(rows int) Spreadsheet {
	return Spreadsheet{m: make(map[string]int)}
}

func (ss *Spreadsheet) SetCell(cell string, value int) {
	ss.m[cell] = value
}

func (ss *Spreadsheet) ResetCell(cell string) {
	delete(ss.m, cell)
}

func (ss *Spreadsheet) Parse(s string) int {
	if unicode.IsLetter(rune(s[0])) {
		return ss.m[s]
	}
	x, _ := strconv.Atoi(s)
	return x
}

func (ss *Spreadsheet) GetValue(formula string) int {
	plusIndex := strings.Index(formula, "+")
	return ss.Parse(formula[1:plusIndex]) + ss.Parse(formula[plusIndex+1:])
}

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */
