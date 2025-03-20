package basic

import (
	"fmt"
)

// Ivec2 is two-component vector of integer numbers.
type Ivec2 [2]int

// Ivec2Num returns vector with all components set to k.
func Ivec2Num(k int) Ivec2 {
	return Ivec2{k, k}
}

// String implements the [fmt.Stringer] interface.
func (v Ivec2) String() string {
	return fmt.Sprintf("(%d, %d)", v[0], v[1])
}

// Prec returns precise version of v.
func (v Ivec2) Prec() Vec2 {
	return Vec2{Float(v[0]), Float(v[1])}
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

// Abs returns |v|.
func (v Ivec2) Abs() Ivec2 {
	return Ivec2{Abs(v[0]), Abs(v[1])}
}

// Neg returns -v.
func (v Ivec2) Neg() Ivec2 {
	return Ivec2{-v[0], -v[1]}
}

// Add returns v + w.
func (v Ivec2) Add(w Ivec2) Ivec2 {
	return Ivec2{v[0] + w[0], v[1] + w[1]}
}

// AddNum returns v + k.
func (v Ivec2) AddNum(k int) Ivec2 {
	return Ivec2{v[0] + k, v[1] + k}
}

// Sub returns v - w.
func (v Ivec2) Sub(w Ivec2) Ivec2 {
	return Ivec2{v[0] - w[0], v[1] - w[1]}
}

// SubNum returns v - k.
func (v Ivec2) SubNum(k int) Ivec2 {
	return Ivec2{v[0] - k, v[1] - k}
}

// Mul returns v * w.
func (v Ivec2) Mul(w Ivec2) Ivec2 {
	return Ivec2{v[0] * w[0], v[1] * w[1]}
}

// MulNum returns v * k.
func (v Ivec2) MulNum(k int) Ivec2 {
	return Ivec2{v[0] * k, v[1] * k}
}

// Div returns v / w.
func (v Ivec2) Div(w Ivec2) Ivec2 {
	return Ivec2{v[0] / w[0], v[1] / w[1]}
}

// DivNum returns v / k.
func (v Ivec2) DivNum(k int) Ivec2 {
	return Ivec2{v[0] / k, v[1] / k}
}

// Mod returns v modulo w (see [Mod]).
func (v Ivec2) Mod(w Ivec2) Ivec2 {
	return Ivec2{Mod(v[0], w[0]), Mod(v[1], w[1])}
}

// ModNum returns v modulo k (see [Mod]).
func (v Ivec2) ModNum(k int) Ivec2 {
	return Ivec2{Mod(v[0], k), Mod(v[1], k)}
}

// L1Norm returns L1 (Manhattan) norm of v.
func (v Ivec2) L1Norm() int {
	return Abs(v[0]) + Abs(v[1])
}
