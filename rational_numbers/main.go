package main

import (
	"fmt"
)

type Rational struct {
	p int
	q int
}

type operation interface {
	plus(that Rational) (Rational)
	minus(that Rational) (Rational)
	times(that Rational) (Rational)
	divides(that Rational) (Rational)
	equals(that Rational) (bool)
	toString() (string)
}

func (r Rational) toString() (string) {
	if r.p == 0 {
		return "0"
	} else if r.p == r.q && r.q == 1 {
		return "1"
	}

	return fmt.Sprintf("%d/%d", r.p, r.q)
}

func (r Rational) equals(t Rational) (bool) {
	return r.p == t.p && r.q == t.q
}

func (r Rational) plus(t Rational) (Rational) {
	n := r.p * t.q + t.p * r.q
	d := r.q * t.q

	if n == d {
		return Rational{1, 1}
	} else {
		return Rational{n, d}
	}
}

func (r Rational) minus(t Rational) (Rational) {
	n := r.p * t.q - t.p * r.q
	d := r.q * t.q

	if n == 0 {
		return Rational{0, 0}
	} else {
		return Rational{n, d}
	}
}

func (r Rational) divides(t Rational) (Rational) {
	n := r.p * t.q 
	d := r.q * t.p

	if n == d {
		return Rational{1, 1}
	} else {
		return Rational{n, d}
	}
}

func (r Rational) times(t Rational) (Rational) {
	n := r.p * t.p
	d := r.q * t.q

	return Rational{n, d}
}

func main() {
	t1 := Rational{1, 2}
	t2 := Rational{1, 2}

	fmt.Printf("t1: %v\n", t1.toString())
	fmt.Printf("t2: %v\n", t2.toString())
	fmt.Printf("t1 == t2: %v\n", t1.equals(t2))
	fmt.Printf("t1 + t2: %s\n", (t1.plus(t2)).toString())
	fmt.Printf("t1 - t2: %s\n", (t1.minus(t2)).toString())
	fmt.Printf("t1 * t2: %s\n", (t1.times(t2)).toString())
	fmt.Printf("t1 / t2: %s\n", (t1.divides(t2)).toString())
}
