package code

import "testing"

func TestDest(t *testing.T) {
	if Dest("M") != "001" {
		t.Error("Dest('M') failed")
	}
}

func TestComp(t *testing.T) {
	if Comp("D") != "0001100" {
		t.Error("Comp('D') failed")
	}
}

func TestJump(t *testing.T) {
	if Jump("JGT") != "001" {
		t.Error("Jump('JGT') failed")
	}
}
