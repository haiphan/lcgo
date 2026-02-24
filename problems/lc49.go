package problems

import "unsafe"

type FastByteTable struct {
	mask   uint32
	keys   [][26]byte
	values []int32 // Using -1 as sentinel
}

func NewFastByteTable(n int) *FastByteTable {
	size := uint32(1)
	for size < uint32(n*2) { // 50% load factor
		size <<= 1
	}

	v := make([]int32, size)
	for i := range v {
		v[i] = -1
	}

	return &FastByteTable{
		mask:   size - 1,
		keys:   make([][26]byte, size),
		values: v,
	}
}

// hash64 processes 26 bytes using 8-byte register jumps
func (t *FastByteTable) hash(key *[26]byte) uint32 {
	var h uint64 = 14695981039346656037 // FNV-1a 64-bit offset
	p := unsafe.Pointer(key)

	// Process 24 bytes (3 chunks of 8)
	h = (h ^ *(*uint64)(p)) * 1099511628211
	h = (h ^ *(*uint64)(unsafe.Add(p, 8))) * 1099511628211
	h = (h ^ *(*uint64)(unsafe.Add(p, 16))) * 1099511628211

	// Process remaining 2 bytes
	h = (h ^ uint64(*(*uint16)(unsafe.Add(p, 24)))) * 1099511628211

	// Fold 64-bit hash into 32-bit for the mask
	return uint32(h ^ (h >> 32))
}

// compare checks 26 bytes using three 64-bit and one 16-bit comparison
func (t *FastByteTable) compare(a, b *[26]byte) bool {
	pa := unsafe.Pointer(a)
	pb := unsafe.Pointer(b)

	// Check 24 bytes in three jumps
	if *(*uint64)(pa) != *(*uint64)(pb) {
		return false
	}
	if *(*uint64)(unsafe.Add(pa, 8)) != *(*uint64)(unsafe.Add(pb, 8)) {
		return false
	}
	if *(*uint64)(unsafe.Add(pa, 16)) != *(*uint64)(unsafe.Add(pb, 16)) {
		return false
	}

	// Check remaining 2 bytes
	return *(*uint16)(unsafe.Add(pa, 24)) == *(*uint16)(unsafe.Add(pb, 24))
}

func (t *FastByteTable) Get(key [26]byte) int {
	idx := t.hash(&key) & t.mask

	for t.values[idx] != -1 {
		if t.compare(&t.keys[idx], &key) {
			return int(t.values[idx])
		}
		idx = (idx + 1) & t.mask
	}
	return 0
}

func (t *FastByteTable) Set(key [26]byte, value int) {
	idx := t.hash(&key) & t.mask

	for t.values[idx] != -1 {
		if t.compare(&t.keys[idx], &key) {
			t.values[idx] = int32(value)
			return
		}
		idx = (idx + 1) & t.mask
	}

	t.keys[idx] = key
	t.values[idx] = int32(value)
}

func (t *FastByteTable) SetIfNotExist(key [26]byte, v int) int {
	h := t.hash(&key)
	idx := h & t.mask

	for t.values[idx] != -1 {
		// If the key already exists, return the current value immediately
		if t.compare(&t.keys[idx], &key) {
			return int(t.values[idx])
		}
		// Collision: move to the next slot
		idx = (idx + 1) & t.mask
	}

	// If we reached here, the key does not exist.
	// We are already at the correct empty slot (idx), so we just fill it.
	t.keys[idx] = key
	t.values[idx] = int32(v)
	return v
}

// tbend

func getIdentity(s string) [26]byte {
	var charCounts [26]byte
	for _, c := range s {
		charCounts[c-'a']++
	}
	return charCounts
}

func groupAnagrams(strs []string) [][]string {
	n := len(strs)
	tb := NewFastByteTable(n)
	res := make([][]string, 0, n)
	cur := 0
	for _, s := range strs {
		h := getIdentity(s)
		hi := tb.SetIfNotExist(h, cur+1)
		if hi == cur+1 {
			cur++
		}
		j := hi - 1
		if len(res) == j {
			res = append(res, []string{s})
		} else {
			res[j] = append(res[j], s)
		}
	}
	return res
}
