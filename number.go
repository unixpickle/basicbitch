package main

import (
	"errors"
	"math/big"
	"strconv"
	"strings"
)

// A Number is an array of bits in little endian bit order.
type Number []bool

// NewNumber creates a binary number using a big integer.
// The resulting Number will have as few bits as possible, such that the MSB is always 1.
func NewNumber(i *big.Int, bitSize int) Number {
	var num big.Int
	num.Set(i)

	if num.Sign() == -1 {
		num.Neg(i)
		return NewNumber(i, bitSize).Neg()
	}

	res := make(Number, 0, bitSize)

	two := big.NewInt(2)
	for i := 0; i < bitSize; i++ {
		var mod big.Int
		num.DivMod(&num, two, &mod)
		res = append(res, mod.Sign() == 1)
	}

	return res
}

// ParseNumber parses a number.
// The string representation may be in decimal, hex ("0x..."), octal ("0..."), or binary ("0b").
// If the number string has a leading "-", the non-negated number will be parsed and then negated.
func ParseNumber(s string, bitSize int) (Number, error) {
	if strings.HasPrefix(s, "-") {
		if n, err := ParseNumber(s[1:], bitSize); err != nil {
			return nil, err
		} else {
			return n.Neg(), nil
		}
	}

	var f big.Float
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0b") || !strings.HasPrefix(s, "0") {
		_, _, err := f.Parse(s, 0)
		if err != nil {
			return nil, err
		}
	} else {
		binString := "0b"
		octalToBin := map[string]string{
			"0": "000",
			"1": "001",
			"2": "010",
			"3": "011",
			"4": "100",
			"5": "101",
			"6": "110",
			"7": "111",
		}
		for _, ch := range s[1:] {
			if bin, ok := octalToBin[string(ch)]; ok {
				binString += bin
			} else {
				return nil, errors.New("unknown octal digit: " + string(ch))
			}
		}
		return ParseNumber(binString, bitSize)
	}
	var i big.Int
	f.Int(&i)
	return NewNumber(&i, bitSize), nil
}

// Sum returns the sum of two equally-sized numbers.
//
// The overflow return value specifies if there was arithmetic overflow.
func Sum(n1, n2 Number) (s Number, overflow bool) {
	if len(n1) != len(n2) {
		panic("numbers must match in size")
	}
	s = make(Number, len(n1))
	lastCarry := 0
	carry := 0
	for i, x := range n1 {
		sum := carry
		if x {
			sum += 1
		}
		if n2[i] {
			sum += 1
		}
		s[i] = ((sum & 1) == 1)
		lastCarry = carry
		carry = sum >> 1
	}
	return s, lastCarry != carry
}

// Neg returns the two's complement of the number.
func (n Number) Neg() Number {
	res := make(Number, len(n))
	for i, x := range n {
		if !x {
			res[i] = true
		}
	}

	carry := false
	for i, x := range res {
		sum := 0
		if carry {
			sum++
		}
		if x {
			sum++
		}
		if i == 0 {
			sum++
		}
		res[i] = ((sum & 1) == 1)
		carry = (sum >= 2)
	}

	return res
}

// SignExtend returns a version of the number which is exactly the given number of bits.
// If the number is too long, it will be truncated rather than sign extended.
func (n Number) SignExtend(bitSize int) Number {
	if len(n) > bitSize {
		return n[:bitSize]
	}
	res := make(Number, bitSize)
	copy(res, n)
	if n[len(n)-1] {
		for i := len(n); i < bitSize; i++ {
			res[i] = true
		}
	}
	return res
}

// Equal returns true if both numbers are the same size and contain the same value.
func (n Number) Equal(n1 Number) bool {
	if len(n) != len(n1) {
		return false
	}
	for i, x := range n {
		if n1[i] != x {
			return false
		}
	}
	return true
}

// BigIntSigned returns a big integer representation of this signed number.
func (n Number) BigIntSigned() *big.Int {
	if n[len(n)-1] {
		res := n.Neg().BigIntSigned()
		res.Neg(res)
		return res
	}
	var r big.Int
	place := big.NewInt(1)
	two := big.NewInt(2)
	for _, x := range n {
		if x {
			r.Add(&r, place)
		}
		place.Mul(place, two)
	}
	return &r
}

// BigIntUnsigned returns a big integer representation of this unsigned number.
func (n Number) BigIntUnsigned() *big.Int {
	var r big.Int
	place := big.NewInt(1)
	two := big.NewInt(2)
	for _, x := range n {
		if x {
			r.Add(&r, place)
		}
		place.Mul(place, two)
	}
	return &r
}

// BaseStringSigned returns a string representation of this number in the given base.
// The receiver is treated as a signed number.
// Allowed bases are 2, 8, 10, and 16.
func (n Number) BaseStringSigned(b int) string {
	res := n.BigIntSigned().Text(b)
	prefix, ok := map[int]string{2: "0b", 8: "0", 10: "", 16: "0x"}[b]
	if !ok {
		panic("unknown base: " + strconv.Itoa(b))
	}
	if strings.HasPrefix(res, "-") {
		return res[:1] + prefix + res[1:]
	} else {
		return prefix + res
	}
}

// BaseStringUnsigned returns a string representation of this number in the given base.
// The receiver is treated as a signed number.
// Allowed bases are 2, 8, 10, and 16.
func (n Number) BaseStringUnsigned(b int) string {
	res := n.BigIntUnsigned().Text(b)
	prefix, ok := map[int]string{2: "0b", 8: "0", 10: "", 16: "0x"}[b]
	if !ok {
		panic("unknown base: " + strconv.Itoa(b))
	}
	return prefix + res
}
