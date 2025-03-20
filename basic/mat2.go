package basic

import (
	"fmt"
	"math"
)

// Mat2 is 2x2 column-major matrix of floating-point numbers.
type Mat2 [2]Vec2

// Mat2Num returns matrix with all components set to k.
func Mat2Num(k Float) Mat2 {
	return Mat2{Vec2Num(k), Vec2Num(k)}
}

// Mat2Diag returns matrix with diagonal components set from v and all other ones set to 0.
//
// Mat2Diag result can be interpreted as scaling matrix.
func Mat2Diag(v Vec2) Mat2 {
	return Mat2{{v[0], 0}, {0, v[1]}}
}

// Mat2DiagNum returns matrix with diagonal components set to k and all other ones set to 0.
//
// Mat2DiagNum result for k = 1 is identity matrix.
//
// Mat2DiagNum result can be interpreted as isotropic scaling matrix.
func Mat2DiagNum(k Float) Mat2 {
	return Mat2{{k, 0}, {0, k}}
}

// Mat2Rot returns matrix representing rotation by r radians counterclockwise.
func Mat2Rot(r Float) Mat2 {
	sin, cos := math.Sincos(r)

	return Mat2{{cos, sin}, {-sin, cos}}
}

// String implements the [fmt.Stringer] interface.
func (m Mat2) String() string {
	return fmt.Sprintf("(%s, %s)", m[0], m[1])
}

// IsZero reports whether m is zero.
func (m Mat2) IsZero() bool {
	return m[0].IsZero() && m[1].IsZero()
}

// Abs returns |m|.
func (m Mat2) Abs() Mat2 {
	return Mat2{m[0].Abs(), m[1].Abs()}
}

// Neg returns -m.
func (m Mat2) Neg() Mat2 {
	return Mat2{m[0].Neg(), m[1].Neg()}
}

// Add returns m + n.
func (m Mat2) Add(n Mat2) Mat2 {
	return Mat2{m[0].Add(n[0]), m[1].Add(n[1])}
}

// AddNum returns m + k.
func (m Mat2) AddNum(k Float) Mat2 {
	return Mat2{m[0].AddNum(k), m[1].AddNum(k)}
}

// Sub returns m - n.
func (m Mat2) Sub(n Mat2) Mat2 {
	return Mat2{m[0].Sub(n[0]), m[1].Sub(n[1])}
}

// SubNum returns m - k.
func (m Mat2) SubNum(k Float) Mat2 {
	return Mat2{m[0].SubNum(k), m[1].SubNum(k)}
}

// Mul returns m * n (NOT linear algebra operation).
func (m Mat2) Mul(n Mat2) Mat2 {
	return Mat2{m[0].Mul(n[0]), m[1].Mul(n[1])}
}

// MulNum returns m * k.
func (m Mat2) MulNum(k Float) Mat2 {
	return Mat2{m[0].MulNum(k), m[1].MulNum(k)}
}

// Div returns m / n.
func (m Mat2) Div(n Mat2) Mat2 {
	return Mat2{m[0].Div(n[0]), m[1].Div(n[1])}
}

// DivNum returns m / k.
func (m Mat2) DivNum(k Float) Mat2 {
	return Mat2{m[0].DivNum(k), m[1].DivNum(k)}
}

// Mod returns m modulo n (see [Mod]).
func (m Mat2) Mod(n Mat2) Mat2 {
	return Mat2{m[0].Mod(n[0]), m[1].Mod(n[1])}
}

// ModNum returns m modulo k (seed [Mod]).
func (m Mat2) ModNum(k Float) Mat2 {
	return Mat2{m[0].ModNum(k), m[1].ModNum(k)}
}

// Adj returns adjugate of m.
func (m Mat2) Adj() Mat2 {
	return Mat2{{m[1][1], -m[0][1]}, {-m[1][0], m[0][0]}}
}

// Det returns determinant of m.
func (m Mat2) Det() Float {
	return m[0][0]*m[1][1] - m[1][0]*m[0][1]
}

// Inv returns inverse of m.
func (m Mat2) Inv() Mat2 {
	return m.Adj().MulNum(1 / m.Det())
}

// Times returns m times n (linear algebra operation).
func (m Mat2) Times(n Mat2) Mat2 {
	return Mat2{m.TimesVec(n[0]), m.TimesVec(n[1])}
}

// TimesVec returns m times v (linear algebra operation).
func (m Mat2) TimesVec(v Vec2) Vec2 {
	return Vec2{m[0][0]*v[0] + m[1][0]*v[1], m[0][1]*v[0] + m[1][1]*v[1]}
}

// Tran returns transpose of m.
func (m Mat2) Tran() Mat2 {
	return Mat2{{m[0][0], m[1][0]}, {m[0][1], m[1][1]}}
}
