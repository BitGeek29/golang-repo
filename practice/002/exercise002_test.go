package main

import (
	"testing"
)

func TestEx002(t *testing.T) {

	want := 0
	got, err := Ex002(-10)
	if err == nil {
		t.Errorf("Ex002() = %v, want %v", got, want)
	}

	want = 10
	got, err = Ex002(0)
	if err == nil && got != uint64(want) {
		t.Errorf("Ex002() = %v, want %v", got, want)
	}

	want = 40320
	got, err = Ex002(8)
	if err == nil && got != uint64(want) {
		t.Errorf("Ex002 = %v, wnat %v", got, want)
	}
}
