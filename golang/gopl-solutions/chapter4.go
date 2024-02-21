package main

import "fmt"

func _swap(x, y int) (int, int) {
	return y, x
}

func _reverse(s []int) {
	// reverse [x, y)
	slen := len(s)
	for i := 0; i < slen/2; i++ {
		s[i], s[slen-i-1] = _swap(s[i], s[slen-i-1])
	}
}

func ex4_3(pArr *[5]int) {
	for i := 0; i < 5/2; i++ {
		pArr[i], pArr[4-i] = _swap(pArr[i], pArr[4-i])
	}
}

func ex4_4(s []int, k int) {
	// related: leetcode-189
	slen := len(s)
	k = k % slen
	_reverse(s[:slen-k])
	_reverse(s[slen-k:])
	_reverse(s)
}

func TestChapter4() {
	// ex4_3
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("[ex4_3] array before reverse: %v\n", arr)
	ex4_3(&arr)
	fmt.Printf("[ex4_3] array after reverse: %v\n", arr)

	// ex4_4
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("[ex4_4] slice before rotate: %v\n", s)
	ex4_4(s, 2)
	fmt.Printf("[ex4_4] array after rotate: %v\n", s)
}
