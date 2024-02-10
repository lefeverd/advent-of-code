package main

import "testing"

func TestPart2(t *testing.T) {
	const want = 281
	if got, _ := ComputePart2("part2_test_input.txt"); got != want {
		t.Errorf("got = %d, want %d", got, want)
	}
}
