package basic

import (
	"fmt"
	"math"
)

// Mat2 is 2x2 column-major matrix of floating-point numbers.
type Mat2 [2]Vec2

// Mat2DiagAll returns matrix with diagonal components set to k and all other ones set to 0.
//
// Mat2DiagAll result for k = 1 is identity matrix.
//
// Mat2DiagAll result can be interpreted as isotropic scaling matrix.
func Mat2DiagAll(k Float) Mat2 {
	return Mat2{{k, 0}, {0, k}}
}

// Mat2Diag returns matrix with diagonal components set from v and all other ones set to 0.
//
// Mat2Diag result can be interpreted as scaling matrix.
func Mat2Diag(v Vec2) Mat2 {
	return Mat2{{v[0], 0}, {0, v[1]}}
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

// Abs returns absolute value of m.
func (m Mat2) Abs() Mat2 {
	return Mat2{m[0].Abs(), m[1].Abs()}
}

// Neg returns m with inverted sign.
func (m Mat2) Neg() Mat2 {
	return Mat2{m[0].Neg(), m[1].Neg()}
}

// Add returns sum of m and n.
func (m Mat2) Add(n Mat2) Mat2 {
	return Mat2{m[0].Add(n[0]), m[1].Add(n[1])}
}

// Sub returns difference of m and n.
func (m Mat2) Sub(n Mat2) Mat2 {
	return Mat2{m[0].Sub(n[0]), m[1].Sub(n[1])}
}

// MulAll returns product of m and k.
func (m Mat2) MulAll(k Float) Mat2 {
	return Mat2{m[0].MulAll(k), m[1].MulAll(k)}
}

// MulVec returns product of m and column v (linear algebra).
func (m Mat2) MulVec(v Vec2) Vec2 {
	return Vec2{m[0][0]*v[0] + m[1][0]*v[1], m[0][1]*v[0] + m[1][1]*v[1]}
}

// Mul returns product of m and n (linear algebra).
func (m Mat2) Mul(n Mat2) Mat2 {
	return Mat2{m.MulVec(n[0]), m.MulVec(n[1])}
}

// DivAll returns quotient of m and k.
func (m Mat2) DivAll(k Float) Mat2 {
	return Mat2{m[0].DivAll(k), m[1].DivAll(k)}
}

// ModAll returns m modulo k.
func (m Mat2) ModAll(k Float) Mat2 {
	return Mat2{m[0].ModAll(k), m[1].ModAll(k)}
}

// Det returns determinant of m.
func (m Mat2) Det() Float {
	return m[0][0]*m[1][1] - m[1][0]*m[0][1]
}

// Tran returns transpose of m.
func (m Mat2) Tran() Mat2 {
	return Mat2{{m[0][0], m[1][0]}, {m[0][1], m[1][1]}}
}

// Adj returns adjugate of m.
func (m Mat2) Adj() Mat2 {
	return Mat2{{m[1][1], -m[0][1]}, {-m[1][0], m[0][0]}}
}

// Inv returns inverse of m.
func (m Mat2) Inv() Mat2 {
	return m.Adj().MulAll(1 / m.Det())
}
