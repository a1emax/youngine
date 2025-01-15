package basic

import (
	"fmt"
	"math"
)

// Vec2 is two-component vector of floating-point numbers.
type Vec2 [2]Float

// String implements the [fmt.Stringer] interface.
func (v Vec2) String() string {
	return fmt.Sprintf("(%g, %g)", v[0], v[1])
}

// IsZero reports whether v is zero.
func (v Vec2) IsZero() bool {
	return v[0] == 0 && v[1] == 0
}

// X returns x-component of v.
func (v Vec2) X() Float {
	return v[0]
}

// Y returns y-component of v.
func (v Vec2) Y() Float {
	return v[1]
}

// Abs returns absolute value of v.
func (v Vec2) Abs() Vec2 {
	return Vec2{math.Abs(v[0]), math.Abs(v[1])}
}

// Neg returns v with inverted sign.
func (v Vec2) Neg() Vec2 {
	return Vec2{-v[0], -v[1]}
}

// Add returns sum of v and w.
func (v Vec2) Add(w Vec2) Vec2 {
	return Vec2{v[0] + w[0], v[1] + w[1]}
}

// Sub returns difference of v and w.
func (v Vec2) Sub(w Vec2) Vec2 {
	return Vec2{v[0] - w[0], v[1] - w[1]}
}

// MulAll returns product of v and k.
func (v Vec2) MulAll(k Float) Vec2 {
	return Vec2{v[0] * k, v[1] * k}
}

// Mul returns product of v and w.
func (v Vec2) Mul(w Vec2) Vec2 {
	return Vec2{v[0] * w[0], v[1] * w[1]}
}

// MulMat returns product of row v and m (linear algebra).
func (v Vec2) MulMat(m Mat2) Vec2 {
	return Vec2{v[0]*m[0][0] + v[1]*m[0][1], v[0]*m[1][0] + v[1]*m[1][1]}
}

// DivAll returns quotient of v and k.
func (v Vec2) DivAll(k Float) Vec2 {
	return Vec2{v[0] / k, v[1] / k}
}

// Div returns quotient of v and w.
func (v Vec2) Div(w Vec2) Vec2 {
	return Vec2{v[0] / w[0], v[1] / w[1]}
}

// ModAll returns v modulo k.
func (v Vec2) ModAll(k Float) Vec2 {
	return Vec2{Mod(v[0], k), Mod(v[1], k)}
}

// Mod returns v modulo w.
func (v Vec2) Mod(w Vec2) Vec2 {
	return Vec2{Mod(v[0], w[0]), Mod(v[1], w[1])}
}

// Mag returns magnitude of v.
func (v Vec2) Mag() Float {
	return math.Sqrt(v.MagSqr())
}

// MagSqr returns magnitude of v squared.
func (v Vec2) MagSqr() Float {
	return v[0]*v[0] + v[1]*v[1]
}

// Normalize returns unit vector in the same direction as v.
func (v Vec2) Normalize() Vec2 {
	return v.MulAll(1 / v.Mag())
}

// Dot returns dot product of v and w.
func (v Vec2) Dot(w Vec2) Float {
	return v[0]*w[0] + v[1]*w[1]
}

// CrossZ returns x- and y-components of cross product of (v.x, v.y, 0) and (0, 0, w.z).
func (v Vec2) CrossZ(wZ Float) Vec2 {
	// v X w = v1*w2 - v2*w1 = v.y*w.z
	//         v2*w0 - v0*w2   -v.x*w.z
	//		   v0*w1 - v1*w0   0
	return Vec2{v[1] * wZ, v[0] * -wZ}
}

// Cross returns z-component of cross product of (v.x, v.y, 0) and (w.x, w.y, 0).
func (v Vec2) Cross(w Vec2) Float {
	// v X w = v1*w2 - v2*w1 = 0
	//         v2*w0 - v0*w2   0
	//		   v0*w1 - v1*w0   v.x*w.y - v.y*w.x
	return v[0]*w[1] - v[1]*w[0]
}
