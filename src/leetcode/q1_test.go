package main

import "testing"

func TestQ1(t *testing.T) {
	input := []int{2, 7, 11, 15}
	got := TwoSum(input, 9)
	want := []int{0, 1}

	if got[0] != want[0] || got[1] != want[1] {
		t.Errorf("got %q want %q", got, want)
	}
}
