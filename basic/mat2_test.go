package basic

import (
	"math"
	"testing"
)

func TestMat2Num(t *testing.T) {
	type args struct {
		k Float
	}
	type want struct {
		result Mat2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "1",
			args: args{
				k: 1,
			},
			want: want{
				result: Mat2{{1, 1}, {1, 1}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Mat2Num(tt.args.k)
			if result != tt.want.result {
				t.Fatalf("(%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestMat2Diag(t *testing.T) {
	type args struct {
		v Vec2
	}
	type want struct {
		result Mat2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1;2)",
			args: args{
				v: Vec2{1, 2},
			},
			want: want{
				result: Mat2{{1, 0}, {0, 2}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Mat2Diag(tt.args.v)
			if result != tt.want.result {
				t.Fatalf("(%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestMat2DiagNum(t *testing.T) {
	type args struct {
		k Float
	}
	type want struct {
		result Mat2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "1",
			args: args{
				k: 1,
			},
			want: want{
				result: Mat2{{1, 0}, {0, 1}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Mat2DiagNum(tt.args.k)
			if result != tt.want.result {
				t.Fatalf("(%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestMat2Rot(t *testing.T) {
	type args struct {
		r Float
	}
	type want struct {
		result Mat2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "pi/2",
			args: args{
				r: math.Pi / 2,
			},
			want: want{
				result: Mat2{
					{math.Cos(math.Pi / 2), math.Sin(math.Pi / 2)},
					{-math.Sin(math.Pi / 2), math.Cos(math.Pi / 2)},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Mat2Rot(tt.args.r)
			if result != tt.want.result {
				t.Fatalf("(%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestMat2_IsZero(t *testing.T) {
	type args struct {
		m Mat2
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
				m: Mat2{},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "((0,0),(0,0))",
			args: args{
				m: Mat2{{0, 0}, {0, 0}},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "((1,0),(0,0))",
			args: args{
				m: Mat2{{1, 0}, {0, 0}},
			},
			want: want{
				result: false,
			},
		},
		{
			name: "((0,0),(1,0))",
			args: args{
				m: Mat2{{0, 0}, {1, 0}},
			},
			want: want{
				result: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.m.IsZero()
			if result != tt.want.result {
				t.Fatalf("(%t expected, got %t", tt.want.result, result)
			}
		})
	}
}

func TestMat2_Abs(t *testing.T) {
	type args struct {
		m Mat2
	}
	type want struct {
		result Mat2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "((1,-2),(-3,4))",
			args: args{
				m: Mat2{{1, -2}, {-3, 4}},
			},
			want: want{
				result: Mat2{{1, 2}, {3, 4}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.m.Abs()
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestMat2_Neg(t *testing.T) {
	type args struct {
		m Mat2
	}
	type want struct {
		result Mat2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "((1,-2),(-3,4))",
			args: args{
				m: Mat2{{1, -2}, {-3, 4}},
			},
			want: want{
				result: Mat2{{-1, 2}, {3, -4}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.m.Neg()
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestMat2_Add(t *testing.T) {
	type args struct {
		m Mat2
		n Mat2
	}
	type want struct {
		result Mat2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "((1,2),(3,4));((10,-20),(-30,40))",
			args: args{
				m: Mat2{{1, 2}, {3, 4}},
				n: Mat2{{10, -20}, {-30, 40}},
			},
			want: want{
				result: Mat2{{11, -18}, {-27, 44}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.m.Add(tt.args.n)
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestMat2_AddNum(t *testing.T) {
	type args struct {
		m Mat2
		k Float
	}
	type want struct {
		result Mat2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "((1,2),(3,4));-10",
			args: args{
				m: Mat2{{1, 2}, {3, 4}},
				k: -10,
			},
			want: want{
				result: Mat2{{-9, -8}, {-7, -6}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.m.AddNum(tt.args.k)
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestMat2_Sub(t *testing.T) {
	type args struct {
		m Mat2
		n Mat2
	}
	type want struct {
		result Mat2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "((1,2),(3,4));((10,-20),(-30,40))",
			args: args{
				m: Mat2{{1, 2}, {3, 4}},
				n: Mat2{{10, -20}, {-30, 40}},
			},
			want: want{
				result: Mat2{{-9, 22}, {33, -36}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.m.Sub(tt.args.n)
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestMat2_SubNum(t *testing.T) {
	type args struct {
		m Mat2
		k Float
	}
	type want struct {
		result Mat2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "((1,2),(3,4));-10",
			args: args{
				m: Mat2{{1, 2}, {3, 4}},
				k: -10,
			},
			want: want{
				result: Mat2{{11, 12}, {13, 14}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.m.SubNum(tt.args.k)
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestMat2_Mul(t *testing.T) {
	type args struct {
		m Mat2
		n Mat2
	}
	type want struct {
		result Mat2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "((1,2),(3,4));((10,-20),(-30,40))",
			args: args{
				m: Mat2{{1, 2}, {3, 4}},
				n: Mat2{{10, -20}, {-30, 40}},
			},
			want: want{
				result: Mat2{{10, -40}, {-90, 160}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.m.Mul(tt.args.n)
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestMat2_MulNum(t *testing.T) {
	type args struct {
		m Mat2
		k Float
	}
	type want struct {
		result Mat2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "((1,2),(3,4));5",
			args: args{
				m: Mat2{{1, 2}, {3, 4}},
				k: 5,
			},
			want: want{
				result: Mat2{{5, 10}, {15, 20}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.m.MulNum(tt.args.k)
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestMat2_Div(t *testing.T) {
	type args struct {
		m Mat2
		n Mat2
	}
	type want struct {
		result Mat2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "((2,8),(18,32));((2,4),(6,8))",
			args: args{
				m: Mat2{{2, 8}, {18, 32}},
				n: Mat2{{2, 4}, {6, 8}},
			},
			want: want{
				result: Mat2{{1, 2}, {3, 4}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.m.Div(tt.args.n)
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestMat2_DivNum(t *testing.T) {
	type args struct {
		m Mat2
		k Float
	}
	type want struct {
		result Mat2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "((2,4),(6,8));2",
			args: args{
				m: Mat2{{2, 4}, {6, 8}},
				k: 2,
			},
			want: want{
				result: Mat2{{1, 2}, {3, 4}},
			},
		},
		{
			name: "((1,2),(3,4));0",
			args: args{
				m: Mat2{{1, 2}, {3, 4}},
				k: 0,
			},
			want: want{
				result: Mat2{{math.Inf(1), math.Inf(1)}, {math.Inf(1), math.Inf(1)}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.m.DivNum(tt.args.k)
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestMat2_Mod(t *testing.T) {
	type args struct {
		m Mat2
		n Mat2
	}
	type want struct {
		result Mat2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "((10,18),(23,28));((3,4),(5,6))",
			args: args{
				m: Mat2{{10, 18}, {23, 28}},
				n: Mat2{{3, 4}, {5, 6}},
			},
			want: want{
				result: Mat2{{1, 2}, {3, 4}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.m.Mod(tt.args.n)
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestMat2_ModNum(t *testing.T) {
	type args struct {
		m Mat2
		k Float
	}
	type want struct {
		result Mat2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "((10,11),(20,21));3",
			args: args{
				m: Mat2{{10, 11}, {20, 21}},
				k: 3,
			},
			want: want{
				result: Mat2{{1, 2}, {2, 0}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.m.ModNum(tt.args.k)
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestMat2_Adj(t *testing.T) {
	type args struct {
		m Mat2
	}
	type want struct {
		result Mat2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "((1,2),(3,4))",
			args: args{
				m: Mat2{{1, 2}, {3, 4}},
			},
			want: want{
				result: Mat2{{4, -2}, {-3, 1}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.m.Adj()
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestMat2_Det(t *testing.T) {
	type args struct {
		m Mat2
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
			name: "((1,2),(3,4))",
			args: args{
				m: Mat2{{1, 2}, {3, 4}},
			},
			want: want{
				result: -2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.m.Det()
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestMat2_Inv(t *testing.T) {
	type args struct {
		m Mat2
	}
	type want struct {
		result Mat2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "((-1,1),(1.5,-1))",
			args: args{
				m: Mat2{{-1, 1}, {1.5, -1}},
			},
			want: want{
				result: Mat2{{2, 2}, {3, 2}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.m.Inv()
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestMat2_Times(t *testing.T) {
	type args struct {
		m Mat2
		n Mat2
	}
	type want struct {
		result Mat2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "((1,2),(3,4));((5,6),(7,8))",
			args: args{
				m: Mat2{{1, 2}, {3, 4}},
				n: Mat2{{5, 6}, {7, 8}},
			},
			want: want{
				result: Mat2{{23, 34}, {31, 46}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.m.Times(tt.args.n)
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestMat2_TimesVec(t *testing.T) {
	type args struct {
		m Mat2
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
			name: "((1,2),(3,4));(5,6)",
			args: args{
				m: Mat2{{1, 2}, {3, 4}},
				v: Vec2{5, 6},
			},
			want: want{
				result: Vec2{23, 34},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.m.TimesVec(tt.args.v)
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestMat2_Tran(t *testing.T) {
	type args struct {
		m Mat2
	}
	type want struct {
		result Mat2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "((1,2),(3,4))",
			args: args{
				m: Mat2{{1, 2}, {3, 4}},
			},
			want: want{
				result: Mat2{{1, 3}, {2, 4}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.m.Tran()
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}
