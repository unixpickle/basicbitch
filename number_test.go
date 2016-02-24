package main

import "testing"

func TestParseNumber(t *testing.T) {
	n, err := ParseNumber("10", 4)
	if err != nil {
		t.Error(err)
	} else if !n.Equal(Number{false, true, false, true}) {
		t.Error("failed to parse '10'")
	}

	n, err = ParseNumber("0b1010", 4)
	if err != nil {
		t.Error(err)
	} else if !n.Equal(Number{false, true, false, true}) {
		t.Error("failed to parse '0b1010'")
	}

	n, err = ParseNumber("012", 4)
	if err != nil {
		t.Error(err)
	} else if !n.Equal(Number{false, true, false, true}) {
		t.Error("failed to parse '012'")
	}

	n, err = ParseNumber("0xa", 4)
	if err != nil {
		t.Error(err)
	} else if !n.Equal(Number{false, true, false, true}) {
		t.Error("failed to parse '0xa'")
	}

	_, err = ParseNumber("08", 32)
	if err == nil {
		t.Error("should have failed to parse '08'")
	}

	n, err = ParseNumber("-0b1010", 5)
	if err != nil {
		t.Error(err)
	} else if !n.Equal(Number{false, true, true, false, true}) {
		t.Error("failed to parse '-0b1010'")
	}
}

func TestSum(t *testing.T) {
	n1, err := ParseNumber("100", 16)
	if err != nil {
		t.Fatal(err)
	}
	n2, err := ParseNumber("-30", 16)
	if err != nil {
		t.Fatal(err)
	}
	n3, err := ParseNumber("70", 16)
	if err != nil {
		t.Fatal(err)
	}
	s, overflow := Sum(n1, n2)
	if overflow {
		t.Error("incorrectly detected overflow for 100 - 30.")
	}
	if !n3.Equal(s) {
		t.Error("computed 100 - 30 incorrectly.")
	}

	n1, err = ParseNumber("31", 6)
	if err != nil {
		t.Fatal(err)
	}
	n2, err = ParseNumber("1", 6)
	if err != nil {
		t.Fatal(err)
	}
	n3, err = ParseNumber("-32", 6)
	if err != nil {
		t.Fatal(err)
	}
	s, overflow = Sum(n1, n2)
	if !overflow {
		t.Error("failed to detect overflow for 31 + 1.")
	}
	if !n3.Equal(s) {
		t.Error("computed 31 + 1 incorrectly.")
	}
}
