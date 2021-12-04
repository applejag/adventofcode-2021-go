package main

import "testing"

func TestBits(t *testing.T) {
	b := Bits{0b01011, 5}
	want := []int{0, 1, 0, 1, 1}

	for i := 0; i < b.Size; i++ {
		got := b.Bit(i)
		if got != want[i] {
			t.Errorf("want %05b, got %05b", want[i], got)
		}
	}
}

func TestBitsString(t *testing.T) {
	b := Bits{0b01011, 5}
	want := "01011"

	got := b.String()
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
