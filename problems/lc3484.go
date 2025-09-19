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

func (this *Spreadsheet) SetCell(cell string, value int) {
	this.m[cell] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	delete(this.m, cell)
}

func (this *Spreadsheet) Parse(s string) int {
	if unicode.IsLetter(rune(s[0])) {
		return this.m[s]
	}
	x, _ := strconv.Atoi(s)
	return x
}

func (this *Spreadsheet) GetValue(formula string) int {
	plusIndex := strings.Index(formula, "+")
	return this.Parse(formula[1:plusIndex]) + this.Parse(formula[plusIndex+1:])
}

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */
