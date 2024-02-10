package main

import "testing"

func TestPart1(t *testing.T) {
	const want = 142
	if got, _ := ComputePart1("part1_test_input.txt"); got != want {
		t.Errorf("got = %d, want %d", got, want)
	}
}
