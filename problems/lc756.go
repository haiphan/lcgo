package problems

import "math/bits"

func pyramidTransition(bottom string, allowed []string) bool {
	var bToH [36]int
	n := len(bottom)
	for _, a := range allowed {
		bToH[6*(a[0]-'A')+a[1]-'A'] |= (1 << int(a[2]-'A'))
	}

	// Encode row as integer: each char takes 5 bits (up to 26 letters)
	encodeRow := func(row []byte) int {
		result := 0
		for _, ch := range row {
			result = (result << 5) | int(ch-'A')
		}
		return result
	}

	memo := make(map[int]bool)

	var backtrack func(row, nextRow []byte, i int) bool
	backtrack = func(row, nextRow []byte, i int) bool {
		if len(row) == 1 {
			return true
		}
		if i == len(nextRow) {
			rowCode := encodeRow(nextRow)
			if val, ok := memo[rowCode]; ok {
				return val
			}
			result := backtrack(nextRow, make([]byte, len(nextRow)-1), 0)
			memo[rowCode] = result
			return result
		}

		chars := uint32(bToH[6*(row[i]-'A')+row[i+1]-'A'])
		for chars != 0 {
			j := bits.TrailingZeros32(chars)
			nextRow[i] = byte('A' + j)
			if backtrack(row, nextRow, i+1) {
				return true
			}
			chars &= (chars - 1)
		}
		return false
	}
	return backtrack([]byte(bottom), make([]byte, n-1), 0)
}
