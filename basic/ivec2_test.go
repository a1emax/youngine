package basic

import (
	"testing"
)

func TestIvec2Num(t *testing.T) {
	type args struct {
		k int
	}
	type want struct {
		result Ivec2
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
				result: Ivec2{1, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Ivec2Num(tt.args.k)
			if result != tt.want.result {
				t.Fatalf("(%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestIvec2_Prec(t *testing.T) {
	type args struct {
		v Ivec2
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
			name: "empty_literal",
			args: args{
				v: Ivec2{},
			},
			want: want{
				result: Vec2{},
			},
		},
		{
			name: "(0,0)",
			args: args{
				v: Ivec2{0, 0},
			},
			want: want{
				result: Vec2{0, 0},
			},
		},
		{
			name: "(1,0)",
			args: args{
				v: Ivec2{1, 0},
			},
			want: want{
				result: Vec2{1, 0},
			},
		},
		{
			name: "(0,1)",
			args: args{
				v: Ivec2{0, 1},
			},
			want: want{
				result: Vec2{0, 1},
			},
		},
		{
			name: "(1,1)",
			args: args{
				v: Ivec2{1, 1},
			},
			want: want{
				result: Vec2{1, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.Prec()
			if result != tt.want.result {
				t.Fatalf("(%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestIvec2_IsZero(t *testing.T) {
	type args struct {
		v Ivec2
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
				v: Ivec2{},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(0,0)",
			args: args{
				v: Ivec2{0, 0},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(1,0)",
			args: args{
				v: Ivec2{1, 0},
			},
			want: want{
				result: false,
			},
		},
		{
			name: "(0,1)",
			args: args{
				v: Ivec2{0, 1},
			},
			want: want{
				result: false,
			},
		},
		{
			name: "(1,1)",
			args: args{
				v: Ivec2{1, 1},
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

func TestIvec2_X(t *testing.T) {
	type args struct {
		v Ivec2
	}
	type want struct {
		result int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,2)",
			args: args{
				v: Ivec2{1, 2},
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
				t.Fatalf("(%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestIvec2_Y(t *testing.T) {
	type args struct {
		v Ivec2
	}
	type want struct {
		result int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,2)",
			args: args{
				v: Ivec2{1, 2},
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
				t.Fatalf("(%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestIvec2_Abs(t *testing.T) {
	type args struct {
		v Ivec2
	}
	type want struct {
		result Ivec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,-2)",
			args: args{
				v: Ivec2{1, -2},
			},
			want: want{
				result: Ivec2{1, 2},
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

func TestIvec2_Neg(t *testing.T) {
	type args struct {
		v Ivec2
	}
	type want struct {
		result Ivec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,-2)",
			args: args{
				v: Ivec2{1, -2},
			},
			want: want{
				result: Ivec2{-1, 2},
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

func TestIvec2_Add(t *testing.T) {
	type args struct {
		v Ivec2
		w Ivec2
	}
	type want struct {
		result Ivec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,2);(10,-20)",
			args: args{
				v: Ivec2{1, 2},
				w: Ivec2{10, -20},
			},
			want: want{
				result: Ivec2{11, -18},
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

func TestIvec2_AddNum(t *testing.T) {
	type args struct {
		v Ivec2
		k int
	}
	type want struct {
		result Ivec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,2);-10",
			args: args{
				v: Ivec2{1, 2},
				k: -10,
			},
			want: want{
				result: Ivec2{-9, -8},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.AddNum(tt.args.k)
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestIvec2_Sub(t *testing.T) {
	type args struct {
		v Ivec2
		w Ivec2
	}
	type want struct {
		result Ivec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,2);(10,-20)",
			args: args{
				v: Ivec2{1, 2},
				w: Ivec2{10, -20},
			},
			want: want{
				result: Ivec2{-9, 22},
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

func TestIvec2_SubNum(t *testing.T) {
	type args struct {
		v Ivec2
		k int
	}
	type want struct {
		result Ivec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,2);-10",
			args: args{
				v: Ivec2{1, 2},
				k: -10,
			},
			want: want{
				result: Ivec2{11, 12},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.SubNum(tt.args.k)
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestIvec2_Mul(t *testing.T) {
	type args struct {
		v Ivec2
		w Ivec2
	}
	type want struct {
		result Ivec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,2);(10;-20)",
			args: args{
				v: Ivec2{1, 2},
				w: Ivec2{10, -20},
			},
			want: want{
				result: Ivec2{10, -40},
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

func TestIvec2_MulNum(t *testing.T) {
	type args struct {
		v Ivec2
		k int
	}
	type want struct {
		result Ivec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,2);-10",
			args: args{
				v: Ivec2{1, 2},
				k: -10,
			},
			want: want{
				result: Ivec2{-10, -20},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.MulNum(tt.args.k)
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestIvec2_Div(t *testing.T) {
	type args struct {
		v Ivec2
		w Ivec2
	}
	type want struct {
		result Ivec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(2,8);(2,4)",
			args: args{
				v: Ivec2{2, 8},
				w: Ivec2{2, 4},
			},
			want: want{
				result: Ivec2{1, 2},
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

func TestIvec2_DivNum(t *testing.T) {
	type args struct {
		v Ivec2
		k int
	}
	type want struct {
		result Ivec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(2,4);2",
			args: args{
				v: Ivec2{2, 4},
				k: 2,
			},
			want: want{
				result: Ivec2{1, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.DivNum(tt.args.k)
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestIvec2_Mod(t *testing.T) {
	type args struct {
		v Ivec2
		w Ivec2
	}
	type want struct {
		result Ivec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(10,11);(3,4)",
			args: args{
				v: Ivec2{10, 11},
				w: Ivec2{3, 4},
			},
			want: want{
				result: Ivec2{1, 3},
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

func TestIvec2_ModNum(t *testing.T) {
	type args struct {
		v Ivec2
		k int
	}
	type want struct {
		result Ivec2
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(10,11);3",
			args: args{
				v: Ivec2{10, 11},
				k: 3,
			},
			want: want{
				result: Ivec2{1, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.ModNum(tt.args.k)
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestIvec2_L1Norm(t *testing.T) {
	type args struct {
		v Ivec2
	}
	type want struct {
		result int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(-3,4)",
			args: args{
				v: Ivec2{-3, 4},
			},
			want: want{
				result: 7,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.v.L1Norm()
			if result != tt.want.result {
				t.Fatalf("(%d expected, got %d", tt.want.result, result)
			}
		})
	}
}
