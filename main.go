package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	op1, op2, op := ParseArgs()

	var answer Number
	var overflow bool
	switch op {
	case "+":
		answer, overflow = Sum(op1, op2)
	case "-":
		answer, overflow = Sum(op1, op2.Neg())
	default:
		fmt.Fprintln(os.Stderr, "invalid operation:", op)
		os.Exit(1)
	}

	maxLens := []int{0, 0}
	for _, base := range []int{2, 8, 10, 16} {
		ans := []string{answer.BaseStringSigned(base), answer.BaseStringUnsigned(base)}
		for i, x := range ans {
			if len(x) > maxLens[i] {
				maxLens[i] = len(x)
			}
		}
	}
	fmt.Println("")
	for _, base := range []int{2, 8, 10, 16} {
		fmt.Println("Base", padStr(strconv.Itoa(base), 2), "|",
			padStr(answer.BaseStringSigned(base), maxLens[0]),
			"(signed) |", padStr(answer.BaseStringUnsigned(base), maxLens[1]),
			"(unsigned)")
	}
	fmt.Println("\nOverflow:", overflow, "\n")
}

func padStr(s string, desiredLen int) string {
	for len(s) < desiredLen {
		s += " "
	}
	return s
}
