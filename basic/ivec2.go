package basic

import (
	"fmt"
)

// Ivec2 is two-component vector of integer numbers.
type Ivec2 [2]int

// String implements the [fmt.Stringer] interface.
func (v Ivec2) String() string {
	return fmt.Sprintf("(%d, %d)", v[0], v[1])
}

// IsZero reports whether v is zero.
func (v Ivec2) IsZero() bool {
	return v[0] == 0 && v[1] == 0
}

// X returns x-component of v.
func (v Ivec2) X() int {
	return v[0]
}

// Y returns y-component of v.
func (v Ivec2) Y() int {
	return v[1]
}

// Abs returns absolute value of v.
func (v Ivec2) Abs() Ivec2 {
	return Ivec2{Abs(v[0]), Abs(v[1])}
}

// Neg returns v with inverted sign.
func (v Ivec2) Neg() Ivec2 {
	return Ivec2{-v[0], -v[1]}
}

// Add returns sum of v and w.
func (v Ivec2) Add(w Ivec2) Ivec2 {
	return Ivec2{v[0] + w[0], v[1] + w[1]}
}

// Sub returns difference of v and w.
func (v Ivec2) Sub(w Ivec2) Ivec2 {
	return Ivec2{v[0] - w[0], v[1] - w[1]}
}

// MulAll returns product of v and k.
func (v Ivec2) MulAll(k int) Ivec2 {
	return Ivec2{v[0] * k, v[1] * k}
}

// Mul returns product of v and w.
func (v Ivec2) Mul(w Ivec2) Ivec2 {
	return Ivec2{v[0] * w[0], v[1] * w[1]}
}

// DivAll returns quotient of v and k.
func (v Ivec2) DivAll(k int) Ivec2 {
	return Ivec2{v[0] / k, v[1] / k}
}

// Div returns quotient of v and w.
func (v Ivec2) Div(w Ivec2) Ivec2 {
	return Ivec2{v[0] / w[0], v[1] / w[1]}
}

// ModAll returns v modulo k.
func (v Ivec2) ModAll(k int) Ivec2 {
	return Ivec2{Mod(v[0], k), Mod(v[1], k)}
}

// Mod returns v modulo w.
func (v Ivec2) Mod(w Ivec2) Ivec2 {
	return Ivec2{Mod(v[0], w[0]), Mod(v[1], w[1])}
}
