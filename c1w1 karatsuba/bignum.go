package main

import (
	"strconv"
)


type bignum []int


func fromString(s string) bignum {
	res := make(bignum, 0, len([]rune(s)))
	for _, r := range s {
		if '0' <= r && r <= '9' {
			res = append(res, int(r) - '0')
		}
	}

	return res
}

func fromInt(n int) bignum {
	s := strconv.Itoa(n)
	return fromString(s)
}

func padLeft(n bignum, count int) bignum {
	pad := make(bignum, count)
	return append(pad, n...)
}

func padRight(n bignum, count int) bignum {
	pad := make(bignum, count)
	return append(n, pad...)
}

func sum(n1, n2 bignum) bignum {
	n1, n2 = alignLen(n1, n2)
	res := make(bignum, len(n1), len(n1))

	prev := 0
	for i := len(n1) - 1; i >= 0; i-- {
		d := n1[i] + n2[i] + prev
		if d < 10 {
			res[i] = d
			prev = 0
		} else {
			res[i] = d - 10
			prev = 1
		}
	}

	if prev == 1 {
		res = append([]int{1}, res...)
	}

	return res
}

// expects n1 >= n2
func subtract(n1, n2 bignum) bignum {
	n1, n2 = alignLen(n1, n2)
	res := make(bignum, len(n1), len(n1))

	prev := 0
	for i := len(n1) - 1; i >= 0; i-- {
		d := n1[i] - n2[i] + prev
		if d >= 0 {
			res[i] = d
			prev = 0
		} else {
			res[i] = 10 + d
			prev = -1
		}
	}

	// // todo: support negative result
	// if prev == -1 {
	// 	res = append([]int{1}, res...)
	// }

	return res
}

func mult(n1, n2 bignum) bignum {
	if len(n1) == 1 && len(n2) == 1 {
		return fromInt(n1[0] * n2[0])
	}

	n1, n2 = alignLen(n1, n2)

	// garantee len to be even
	if len(n1) % 2 > 0 {
		n1 = padLeft(n1, 1)
		n2 = padLeft(n2, 1)
	}

	a, b := n1[:len(n1) / 2], n1[len(n1) / 2:]
	c, d := n2[:len(n2) / 2], n2[len(n2) / 2:]

	ac := mult(a, c)
	bd := mult(b, d)

	s3 := mult(sum(a, b), sum(c, d)) // (a + b)(c + d)
	s4 := subtract(subtract(s3, ac), bd) // ad + bc

	return sum(sum(padRight(ac, len(n1)), bd), padRight(s4, len(n1) / 2))
}

func alignLen(n1, n2 bignum) (bignum, bignum) {
	diff := len(n1) - len(n2)
	if diff > 0 {
		n2 = padLeft(n2, diff)
	} else if diff < 0 {
		n1 = padLeft(n1, -diff)
	}

	return n1, n2
}
