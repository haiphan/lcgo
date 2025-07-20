package problems

import (
	"sort"
	"strings"
)

type Trie struct {
	id       string
	desc     string
	children map[string]*Trie
}

func createNode(id string) *Trie {
	return &Trie{id: id, desc: "", children: make(map[string]*Trie, 0)}
}

func insertTrie(root *Trie, path []string) {
	cur := root
	for _, name := range path {
		_, has := cur.children[name]
		if !has {
			cur.children[name] = createNode(name)
		}
		cur = cur.children[name]
	}
}

func deleteDuplicateFolder(paths [][]string) [][]string {
	trie := createNode("/")
	for _, p := range paths {
		insertTrie(trie, p)
	}
	dCount := make(map[string]int)
	res := make([][]string, 0)
	cur := make([]string, 0)
	var buildFolder func(node *Trie)
	var collect func(node *Trie)
	buildFolder = func(node *Trie) {
		if len(node.children) == 0 {
			return
		}
		arr := make([]string, len(node.children))
		i := 0
		for k, c := range node.children {
			buildFolder(c)
			arr[i] = k + "(" + c.desc + ")"
			i++
		}
		sort.Strings(arr)
		d := strings.Join(arr, "")
		node.desc = d
		dCount[d]++
	}

	collect = func(node *Trie) {
		if dCount[node.desc] > 1 {
			return
		}
		if len(cur) > 0 {
			dc := make([]string, len(cur))
			copy(dc, cur)
			res = append(res, dc)
		}
		for k, c := range node.children {
			cur = append(cur, k)
			collect(c)
			cur = cur[:len(cur)-1]
		}
	}

	buildFolder(trie)
	collect(trie)
	return res
}
