package basic

import (
	"math"
	"testing"
)

func TestVec2_IsZero(t *testing.T) {
	type args struct {
		v Vec2
	}
	type want struct {
		result bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "empty_literal",
			args: args{
				v: Vec2{},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(0,0)",
			args: args{
				v: Vec2{0, 0},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(1,0)",
			args: args{
				v: Vec2{1, 0},
			},
			want: want{
				result: false,
			},
		},
		{
			name: "(0,1)",
			args: args{
				v: Vec2{0, 1},
			},
			want: want{
				result: false,
			},
		},
		{
			name: "(1,1)",
			args: args{
				v: Vec2{1, 1},
			},
			want: want{
				result: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.IsZero()
			if result != tt.want.result {
				t.Fatalf("(%t expected, got %t", tt.want.result, result)
			}
		})
	}
}

func TestVec2_X(t *testing.T) {
	type args struct {
		v Vec2
	}
	type want struct {
		result Float
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,2)",
			args: args{
				v: Vec2{1, 2},
			},
			want: want{
				result: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.X()
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestVec2_Y(t *testing.T) {
	type args struct {
		v Vec2
	}
	type want struct {
		result Float
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,2)",
			args: args{
				v: Vec2{1, 2},
			},
			want: want{
				result: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.Y()
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestVec2_Abs(t *testing.T) {
	type args struct {
		v Vec2
	}
	type want struct {
		result Vec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,-2)",
			args: args{
				v: Vec2{1, -2},
			},
			want: want{
				result: Vec2{1, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.Abs()
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestVec2_Neg(t *testing.T) {
	type args struct {
		v Vec2
	}
	type want struct {
		result Vec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,-2)",
			args: args{
				v: Vec2{1, -2},
			},
			want: want{
				result: Vec2{-1, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.Neg()
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestVec2_Add(t *testing.T) {
	type args struct {
		v Vec2
		w Vec2
	}
	type want struct {
		result Vec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,2);(10,-20)",
			args: args{
				v: Vec2{1, 2},
				w: Vec2{10, -20},
			},
			want: want{
				result: Vec2{11, -18},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.Add(tt.args.w)
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestVec2_Sub(t *testing.T) {
	type args struct {
		v Vec2
		w Vec2
	}
	type want struct {
		result Vec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,2);(10,-20)",
			args: args{
				v: Vec2{1, 2},
				w: Vec2{10, -20},
			},
			want: want{
				result: Vec2{-9, 22},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.Sub(tt.args.w)
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestVec2_MulAll(t *testing.T) {
	type args struct {
		v Vec2
		k Float
	}
	type want struct {
		result Vec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,2);-10",
			args: args{
				v: Vec2{1, 2},
				k: -10,
			},
			want: want{
				result: Vec2{-10, -20},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.MulAll(tt.args.k)
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestVec2_Mul(t *testing.T) {
	type args struct {
		v Vec2
		w Vec2
	}
	type want struct {
		result Vec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,2);(10;-20)",
			args: args{
				v: Vec2{1, 2},
				w: Vec2{10, -20},
			},
			want: want{
				result: Vec2{10, -40},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.Mul(tt.args.w)
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestVec2_MulMat(t *testing.T) {
	type args struct {
		v Vec2
		m Mat2
	}
	type want struct {
		result Vec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(5,6);((1,2),(3,4))",
			args: args{
				v: Vec2{5, 6},
				m: Mat2{{1, 2}, {3, 4}},
			},
			want: want{
				result: Vec2{17, 39},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.MulMat(tt.args.m)
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestVec2_DivAll(t *testing.T) {
	type args struct {
		v Vec2
		k Float
	}
	type want struct {
		result Vec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(2,4);2",
			args: args{
				v: Vec2{2, 4},
				k: 2,
			},
			want: want{
				result: Vec2{1, 2},
			},
		},
		{
			name: "(1,2);0",
			args: args{
				v: Vec2{1, 2},
				k: 0,
			},
			want: want{
				result: Vec2{math.Inf(1), math.Inf(1)},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.DivAll(tt.args.k)
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestVec2_Div(t *testing.T) {
	type args struct {
		v Vec2
		w Vec2
	}
	type want struct {
		result Vec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(2,8);(2,4)",
			args: args{
				v: Vec2{2, 8},
				w: Vec2{2, 4},
			},
			want: want{
				result: Vec2{1, 2},
			},
		},
		{
			name: "(1,2);(0,1)",
			args: args{
				v: Vec2{1, 2},
				w: Vec2{0, 1},
			},
			want: want{
				result: Vec2{math.Inf(1), 2},
			},
		},
		{
			name: "(1,2);(1,0)",
			args: args{
				v: Vec2{1, 2},
				w: Vec2{1, 0},
			},
			want: want{
				result: Vec2{1, math.Inf(1)},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.Div(tt.args.w)
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestVec2_ModAll(t *testing.T) {
	type args struct {
		v Vec2
		k Float
	}
	type want struct {
		result Vec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(10,11);3",
			args: args{
				v: Vec2{10, 11},
				k: 3,
			},
			want: want{
				result: Vec2{1, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.ModAll(tt.args.k)
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestVec2_Mod(t *testing.T) {
	type args struct {
		v Vec2
		w Vec2
	}
	type want struct {
		result Vec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(10,11);(3,4)",
			args: args{
				v: Vec2{10, 11},
				w: Vec2{3, 4},
			},
			want: want{
				result: Vec2{1, 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.Mod(tt.args.w)
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestVec2_Mag(t *testing.T) {
	type args struct {
		v Vec2
	}
	type want struct {
		result Float
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(3,4)",
			args: args{
				v: Vec2{3, 4},
			},
			want: want{
				result: 5,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.Mag()
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestVec2_MagSqr(t *testing.T) {
	type args struct {
		v Vec2
	}
	type want struct {
		result Float
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(3,4)",
			args: args{
				v: Vec2{3, 4},
			},
			want: want{
				result: 25,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.MagSqr()
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestVec2_Normalize(t *testing.T) {
	type args struct {
		v Vec2
	}
	type want struct {
		result Vec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(3,4)",
			args: args{
				v: Vec2{3, 4},
			},
			want: want{
				result: Vec2{0.6000000000000001, 0.8},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.Normalize()
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestVec2_Dot(t *testing.T) {
	type args struct {
		v Vec2
		w Vec2
	}
	type want struct {
		result Float
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,2);(10,-20)",
			args: args{
				v: Vec2{1, 2},
				w: Vec2{10, -20},
			},
			want: want{
				result: -30,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.Dot(tt.args.w)
			if result != tt.want.result {
				t.Fatalf("%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestVec2_CrossZ(t *testing.T) {
	type args struct {
		v  Vec2
		wZ Float
	}
	type want struct {
		result Vec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,2);-10",
			args: args{
				v:  Vec2{1, 2},
				wZ: -10,
			},
			want: want{
				result: Vec2{-20, 10},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.CrossZ(tt.args.wZ)
			if result != tt.want.result {
				t.Fatalf("%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestVec2_Cross(t *testing.T) {
	type args struct {
		v Vec2
		w Vec2
	}
	type want struct {
		result Float
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,2);(10,-20)",
			args: args{
				v: Vec2{1, 2},
				w: Vec2{10, -20},
			},
			want: want{
				result: -40,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.Cross(tt.args.w)
			if result != tt.want.result {
				t.Fatalf("%g expected, got %g", tt.want.result, result)
			}
		})
	}
}
