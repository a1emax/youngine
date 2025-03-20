package basic

import (
	"math"
	"reflect"
	"testing"
)

func TestAbs_Int(t *testing.T) {
	type args struct {
		x int
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
			name: "0",
			args: args{
				x: 0,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "1",
			args: args{
				x: 1,
			},
			want: want{
				result: 1,
			},
		},
		{
			name: "-1",
			args: args{
				x: -1,
			},
			want: want{
				result: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Abs(tt.args.x)
			if result != tt.want.result {
				t.Fatalf("%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestAbs_Float(t *testing.T) {
	type args struct {
		x Float
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
			name: "0",
			args: args{
				x: 0,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "1",
			args: args{
				x: 1,
			},
			want: want{
				result: 1,
			},
		},
		{
			name: "-1",
			args: args{
				x: -1,
			},
			want: want{
				result: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Abs(tt.args.x)
			if result != tt.want.result {
				t.Fatalf("%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestClamp__Int(t *testing.T) {
	type args struct {
		x    int
		low  int
		high int
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
			name: "5;[1,10]",
			args: args{
				x:    5,
				low:  1,
				high: 10,
			},
			want: want{
				result: 5,
			},
		},
		{
			name: "0;[1,10]",
			args: args{
				x:    0,
				low:  1,
				high: 10,
			},
			want: want{
				result: 1,
			},
		},
		{
			name: "11;[1,10]",
			args: args{
				x:    11,
				low:  1,
				high: 10,
			},
			want: want{
				result: 10,
			},
		},
		{
			name: "5;[10,1]",
			args: args{
				x:    5,
				low:  10,
				high: 1,
			},
			want: want{
				result: 10,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Clamp(tt.args.x, tt.args.low, tt.args.high)
			if result != tt.want.result {
				t.Fatalf("(%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestClamp__Uint(t *testing.T) {
	type args struct {
		x    uint
		low  uint
		high uint
	}
	type want struct {
		result uint
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "5;[1,10]",
			args: args{
				x:    5,
				low:  1,
				high: 10,
			},
			want: want{
				result: 5,
			},
		},
		{
			name: "0;[1,10]",
			args: args{
				x:    0,
				low:  1,
				high: 10,
			},
			want: want{
				result: 1,
			},
		},
		{
			name: "11;[1,10]",
			args: args{
				x:    11,
				low:  1,
				high: 10,
			},
			want: want{
				result: 10,
			},
		},
		{
			name: "5;[10,1]",
			args: args{
				x:    5,
				low:  10,
				high: 1,
			},
			want: want{
				result: 10,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Clamp(tt.args.x, tt.args.low, tt.args.high)
			if result != tt.want.result {
				t.Fatalf("(%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestClamp__Float(t *testing.T) {
	type args struct {
		x    Float
		low  Float
		high Float
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
			name: "5;[1,10]",
			args: args{
				x:    5,
				low:  1,
				high: 10,
			},
			want: want{
				result: 5,
			},
		},
		{
			name: "0;[1,10]",
			args: args{
				x:    0,
				low:  1,
				high: 10,
			},
			want: want{
				result: 1,
			},
		},
		{
			name: "11;[1,10]",
			args: args{
				x:    11,
				low:  1,
				high: 10,
			},
			want: want{
				result: 10,
			},
		},
		{
			name: "5;[10,1]",
			args: args{
				x:    5,
				low:  10,
				high: 1,
			},
			want: want{
				result: 10,
			},
		},
		{
			name: "5;[-Inf,10]",
			args: args{
				x:    5,
				low:  math.Inf(-1),
				high: 10,
			},
			want: want{
				result: 5,
			},
		},
		{
			name: "5;[1,+Inf]",
			args: args{
				x:    5,
				low:  1,
				high: math.Inf(1),
			},
			want: want{
				result: 5,
			},
		},
		{
			name: "5;[-Inf,+Inf]",
			args: args{
				x:    5,
				low:  math.Inf(-1),
				high: math.Inf(1),
			},
			want: want{
				result: 5,
			},
		},
		{
			name: "11;[-Inf,10]",
			args: args{
				x:    11,
				low:  math.Inf(-1),
				high: 10,
			},
			want: want{
				result: 10,
			},
		},
		{
			name: "1;[10,+Inf]",
			args: args{
				x:    1,
				low:  10,
				high: math.Inf(1),
			},
			want: want{
				result: 10,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Clamp(tt.args.x, tt.args.low, tt.args.high)
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestClamp__String(t *testing.T) {
	type args struct {
		x    string
		low  string
		high string
	}
	type want struct {
		result string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "ab;[aa,ac]",
			args: args{
				x:    "ab",
				low:  "aa",
				high: "ac",
			},
			want: want{
				result: "ab",
			},
		},
		{
			name: "a;[aa,ac]",
			args: args{
				x:    "a",
				low:  "aa",
				high: "ac",
			},
			want: want{
				result: "aa",
			},
		},
		{
			name: "ad;[aa,ac]",
			args: args{
				x:    "ad",
				low:  "aa",
				high: "ac",
			},
			want: want{
				result: "ac",
			},
		},
		{
			name: "ab;[ac,aa]",
			args: args{
				x:    "ab",
				low:  "ac",
				high: "aa",
			},
			want: want{
				result: "ac",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Clamp(tt.args.x, tt.args.low, tt.args.high)
			if result != tt.want.result {
				t.Fatalf("(%q expected, got %q", tt.want.result, result)
			}
		})
	}
}

func TestFloorPoz__Float(t *testing.T) {
	type args struct {
		x Float
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
			name: "0",
			args: args{
				x: 0,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "42",
			args: args{
				x: 42,
			},
			want: want{
				result: 42,
			},
		},
		{
			name: "42.5",
			args: args{
				x: 42.5,
			},
			want: want{
				result: 42,
			},
		},
		{
			name: "-42",
			args: args{
				x: -42,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "+Inf",
			args: args{
				x: math.Inf(1),
			},
			want: want{
				result: math.Inf(1),
			},
		},
		{
			name: "-Inf",
			args: args{
				x: math.Inf(-1),
			},
			want: want{
				result: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FloorPoz(tt.args.x)
			if result != tt.want.result {
				t.Fatalf("%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestFract__Float(t *testing.T) {
	type args struct {
		x Float
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
			name: "0",
			args: args{
				x: 0,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "1",
			args: args{
				x: 1,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "0.123",
			args: args{
				x: 0.123,
			},
			want: want{
				result: 0.123,
			},
		},
		{
			name: "1.123",
			args: args{
				x: 1.123,
			},
			want: want{
				result: 0.123,
			},
		},
		{
			name: "-1",
			args: args{
				x: -1,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "-0.123",
			args: args{
				x: -0.123,
			},
			want: want{
				result: 0.123,
			},
		},
		{
			name: "-1.123",
			args: args{
				x: -1.123,
			},
			want: want{
				result: 0.123,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Fract(tt.args.x)
			if result != tt.want.result {
				t.Fatalf("%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestIsMaxInt__Int(t *testing.T) {
	type args struct {
		x int
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
			name: "0",
			args: args{
				x: 0,
			},
			want: want{
				result: false,
			},
		},
		{
			name: "min",
			args: args{
				x: math.MinInt,
			},
			want: want{
				result: false,
			},
		},
		{
			name: "max",
			args: args{
				x: math.MaxInt,
			},
			want: want{
				result: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsMaxInt(tt.args.x)
			if result != tt.want.result {
				t.Fatalf("%t expected, got %t", tt.want.result, result)
			}
		})
	}
}

func TestIsMaxUint__Uint(t *testing.T) {
	type args struct {
		x uint
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
			name: "0",
			args: args{
				x: 0,
			},
			want: want{
				result: false,
			},
		},
		{
			name: "max",
			args: args{
				x: math.MaxUint,
			},
			want: want{
				result: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsMaxUint(tt.args.x)
			if result != tt.want.result {
				t.Fatalf("%t expected, got %t", tt.want.result, result)
			}
		})
	}
}

func TestIsMinInt__Int(t *testing.T) {
	type args struct {
		x int
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
			name: "0",
			args: args{
				x: 0,
			},
			want: want{
				result: false,
			},
		},
		{
			name: "min",
			args: args{
				x: math.MinInt,
			},
			want: want{
				result: true,
			},
		},
		{
			name: "max",
			args: args{
				x: math.MaxInt,
			},
			want: want{
				result: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsMinInt(tt.args.x)
			if result != tt.want.result {
				t.Fatalf("%t expected, got %t", tt.want.result, result)
			}
		})
	}
}

func TestMod__Int(t *testing.T) {
	type args struct {
		x int
		y int
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
			name: "0;42",
			args: args{
				x: 0,
				y: 42,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "1;42",
			args: args{
				x: 1,
				y: 42,
			},
			want: want{
				result: 1,
			},
		},
		{
			name: "42;42",
			args: args{
				x: 42,
				y: 42,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "43;42",
			args: args{
				x: 43,
				y: 42,
			},
			want: want{
				result: 1,
			},
		},
		{
			name: "-1;42",
			args: args{
				x: -1,
				y: 42,
			},
			want: want{
				result: 41,
			},
		},
		{
			name: "-42;42",
			args: args{
				x: -42,
				y: 42,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "-43;42",
			args: args{
				x: -43,
				y: 42,
			},
			want: want{
				result: 41,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Mod(tt.args.x, tt.args.y)
			if result != tt.want.result {
				t.Fatalf("%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestLoop__Int(t *testing.T) {
	type args struct {
		x    []int
		low  int
		high int
	}
	type want struct {
		result []int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "0",
			args: args{
				x:    []int{-2, -1, 0, 1, 2},
				low:  0,
				high: 0,
			},
			want: want{
				result: []int{0, 0, 0, 0, 0},
			},
		},
		{
			name: "2",
			args: args{
				x:    []int{-2, -1, 0, 1, 2, 3, 4},
				low:  3,
				high: 3,
			},
			want: want{
				result: []int{3, 3, 3, 3, 3, 3, 3},
			},
		},
		{
			name: "3",
			args: args{
				x:    []int{-2, -1, 0, 1, 2, 3, 4, 5},
				low:  3,
				high: 3,
			},
			want: want{
				result: []int{3, 3, 3, 3, 3, 3, 3, 3},
			},
		},
		{
			name: "[0,3]",
			args: args{
				x:    []int{-8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
				low:  0,
				high: 3,
			},
			want: want{
				result: []int{0, 1, 2, 3, 0, 1, 2, 3, 0, 1, 2, 3, 0, 1, 2, 3, 0, 1, 2, 3},
			},
		},
		{
			name: "[3,5]",
			args: args{
				x:    []int{-6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
				low:  3,
				high: 5,
			},
			want: want{
				result: []int{3, 4, 5, 3, 4, 5, 3, 4, 5, 3, 4, 5, 3, 4, 5, 3, 4, 5},
			},
		},
		{
			name: "-2",
			args: args{
				x:    []int{-4, -3, -2, -1, 0, 1, 2},
				low:  -2,
				high: -2,
			},
			want: want{
				result: []int{-2, -2, -2, -2, -2, -2, -2},
			},
		},
		{
			name: "-3",
			args: args{
				x:    []int{-5, -4, -3, -2, -1, 0, 1, 2},
				low:  -3,
				high: -3,
			},
			want: want{
				result: []int{-3, -3, -3, -3, -3, -3, -3, -3},
			},
		},
		{
			name: "[-4,-1]",
			args: args{
				x:    []int{-12, -11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7},
				low:  -4,
				high: -1,
			},
			want: want{
				result: []int{-4, -3, -2, -1, -4, -3, -2, -1, -4, -3, -2, -1, -4, -3, -2, -1, -4, -3, -2, -1},
			},
		},
		{
			name: "[-4,-2]",
			args: args{
				x:    []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7},
				low:  -4,
				high: -2,
			},
			want: want{
				result: []int{-4, -3, -2, -4, -3, -2, -4, -3, -2, -4, -3, -2, -4, -3, -2, -4, -3, -2},
			},
		},
		{
			name: "[-1,1]",
			args: args{
				x:    []int{-7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7},
				low:  -1,
				high: 1,
			},
			want: want{
				result: []int{-1, 0, 1, -1, 0, 1, -1, 0, 1, -1, 0, 1, -1, 0, 1},
			},
		},
		{
			name: "[1,-1]",
			args: args{
				x:    []int{-3, -2, -1, 0, 1, 2, 3},
				low:  1,
				high: -1,
			},
			want: want{
				result: []int{1, 1, 1, 1, 1, 1, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := make([]int, len(tt.args.x))
			for i := 0; i < len(tt.args.x); i++ {
				result[i] = Loop(tt.args.x[i], tt.args.low, tt.args.high)
			}
			if !reflect.DeepEqual(result, tt.want.result) {
				t.Fatalf("(%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestMaxInt__Int(t *testing.T) {
	result := MaxInt[int]()
	if wantResult := math.MaxInt; result != wantResult {
		t.Fatalf("%d expected, got %d", wantResult, result)
	}
}

func TestMaxUint__Uint(t *testing.T) {
	result := MaxUint[uint]()
	if wantResult := uint(math.MaxUint); result != wantResult {
		t.Fatalf("%d expected, got %d", wantResult, result)
	}
}

func TestMinInt__Int(t *testing.T) {
	result := MinInt[int]()
	if wantResult := math.MinInt; result != wantResult {
		t.Fatalf("%d expected, got %d", wantResult, result)
	}
}

func TestMod__Uint(t *testing.T) {
	type args struct {
		x uint
		y uint
	}
	type want struct {
		result uint
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "0;42",
			args: args{
				x: 0,
				y: 42,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "1;42",
			args: args{
				x: 1,
				y: 42,
			},
			want: want{
				result: 1,
			},
		},
		{
			name: "42;42",
			args: args{
				x: 42,
				y: 42,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "43;42",
			args: args{
				x: 43,
				y: 42,
			},
			want: want{
				result: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Mod(tt.args.x, tt.args.y)
			if result != tt.want.result {
				t.Fatalf("%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestMod__Float(t *testing.T) {
	type args struct {
		x Float
		y Float
	}
	type want struct {
		result Float
		isNaN  bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "0;42",
			args: args{
				x: 0,
				y: 42,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "1;42",
			args: args{
				x: 1,
				y: 42,
			},
			want: want{
				result: 1,
			},
		},
		{
			name: "42;42",
			args: args{
				x: 42,
				y: 42,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "43;42",
			args: args{
				x: 43,
				y: 42,
			},
			want: want{
				result: 1,
			},
		},
		{
			name: "-1;42",
			args: args{
				x: -1,
				y: 42,
			},
			want: want{
				result: 41,
			},
		},
		{
			name: "-42;42",
			args: args{
				x: -42,
				y: 42,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "-43;42",
			args: args{
				x: -43,
				y: 42,
			},
			want: want{
				result: 41,
			},
		},
		{
			name: "1;0",
			args: args{
				x: 1,
				y: 0,
			},
			want: want{
				isNaN: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Mod(tt.args.x, tt.args.y)

			if tt.want.isNaN {
				if !math.IsNaN(result) {
					t.Fatalf("NaN expected, got %g", result)
				}

				return
			}

			if result != tt.want.result {
				t.Fatalf("%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestPoz__Int(t *testing.T) {
	type args struct {
		x int
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
			name: "0",
			args: args{
				x: 0,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "42",
			args: args{
				x: 42,
			},
			want: want{
				result: 42,
			},
		},
		{
			name: "-42",
			args: args{
				x: -42,
			},
			want: want{
				result: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Poz(tt.args.x)
			if result != tt.want.result {
				t.Fatalf("%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestPoz__Float(t *testing.T) {
	type args struct {
		x Float
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
			name: "0",
			args: args{
				x: 0,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "42",
			args: args{
				x: 42,
			},
			want: want{
				result: 42,
			},
		},
		{
			name: "-42",
			args: args{
				x: -42,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "+Inf",
			args: args{
				x: math.Inf(1),
			},
			want: want{
				result: math.Inf(1),
			},
		},
		{
			name: "-Inf",
			args: args{
				x: math.Inf(-1),
			},
			want: want{
				result: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Poz(tt.args.x)
			if result != tt.want.result {
				t.Fatalf("%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestRound__Float(t *testing.T) {
	type args struct {
		x Float
		n int
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
			name: "1.234;2",
			args: args{
				x: 1.234,
				n: 2,
			},
			want: want{
				result: 1.23,
			},
		},
		{
			name: "1;2",
			args: args{
				x: 1,
				n: 2,
			},
			want: want{
				result: 1,
			},
		},
		{
			name: "1.234;0",
			args: args{
				x: 1.234,
				n: 0,
			},
			want: want{
				result: 1,
			},
		},
		{
			name: "1.234;-1",
			args: args{
				x: 1.234,
				n: -1,
			},
			want: want{
				result: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Round(tt.args.x, tt.args.n)
			if result != tt.want.result {
				t.Fatalf("%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestSign__Int(t *testing.T) {
	type args struct {
		x int
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
			name: "0",
			args: args{
				x: 0,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "42",
			args: args{
				x: 42,
			},
			want: want{
				result: 1,
			},
		},
		{
			name: "-42",
			args: args{
				x: -42,
			},
			want: want{
				result: -1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sign(tt.args.x)
			if result != tt.want.result {
				t.Fatalf("%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestSign__Float(t *testing.T) {
	type args struct {
		x Float
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
			name: "0",
			args: args{
				x: 0,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "42",
			args: args{
				x: 42,
			},
			want: want{
				result: 1,
			},
		},
		{
			name: "-42",
			args: args{
				x: -42,
			},
			want: want{
				result: -1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sign(tt.args.x)
			if result != tt.want.result {
				t.Fatalf("%g expected, got %g", tt.want.result, result)
			}
		})
	}
}
