package safeint

import (
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	type args struct {
		x int
	}
	type want struct {
		result   int
		overflow bool
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
		{
			name: "overflow_max",
			args: args{
				x: math.MinInt,
			},
			want: want{
				overflow: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := Abs(tt.args.x)
			if !tt.want.overflow {
				if !ok {
					t.Fatalf("unexpected overflow")
				}
				if result != tt.want.result {
					t.Fatalf("%d expected, got %d", tt.want.result, result)
				}
			} else {
				if ok {
					t.Fatalf("overflow expected, got %d", result)
				}
			}
		})
	}
}

func TestNeg(t *testing.T) {
	type args struct {
		x int
	}
	type want struct {
		result   int
		overflow bool
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
				result: -1,
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
		{
			name: "overflow_max",
			args: args{
				x: math.MinInt,
			},
			want: want{
				overflow: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := Neg(tt.args.x)
			if !tt.want.overflow {
				if !ok {
					t.Fatalf("unexpected overflow")
				}
				if result != tt.want.result {
					t.Fatalf("%d expected, got %d", tt.want.result, result)
				}
			} else {
				if ok {
					t.Fatalf("overflow expected, got %d", result)
				}
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	type want struct {
		result   int
		overflow bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "1;2",
			args: args{
				x: 1,
				y: 2,
			},
			want: want{
				result: 3,
			},
		},
		{
			name: "overflow_max",
			args: args{
				x: math.MaxInt - 2,
				y: 3,
			},
			want: want{
				overflow: true,
			},
		},
		{
			name: "overflow_min",
			args: args{
				x: math.MinInt + 2,
				y: -3,
			},
			want: want{
				overflow: true,
			},
		},
		{
			name: "min;min",
			args: args{
				x: math.MinInt,
				y: math.MinInt,
			},
			want: want{
				overflow: true,
			},
		},
		{
			name: "min;max",
			args: args{
				x: math.MinInt,
				y: math.MaxInt,
			},
			want: want{
				result: -1,
			},
		},
		{
			name: "max;min",
			args: args{
				x: math.MaxInt,
				y: math.MinInt,
			},
			want: want{
				result: -1,
			},
		},
		{
			name: "max;max",
			args: args{
				x: math.MaxInt,
				y: math.MaxInt,
			},
			want: want{
				overflow: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := Add(tt.args.x, tt.args.y)
			if !tt.want.overflow {
				if !ok {
					t.Fatalf("unexpected overflow")
				}
				if result != tt.want.result {
					t.Fatalf("%d expected, got %d", tt.want.result, result)
				}
			} else {
				if ok {
					t.Fatalf("overflow expected, got %d", result)
				}
			}
		})
	}
}

func TestSub(t *testing.T) {
	type args struct {
		x int
		y int
	}
	type want struct {
		result   int
		overflow bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "1;2",
			args: args{
				x: 1,
				y: 2,
			},
			want: want{
				result: -1,
			},
		},
		{
			name: "overflow_max",
			args: args{
				x: math.MaxInt - 2,
				y: -3,
			},
			want: want{
				overflow: true,
			},
		},
		{
			name: "overflow_min",
			args: args{
				x: math.MinInt + 2,
				y: 3,
			},
			want: want{
				overflow: true,
			},
		},
		{
			name: "0;min",
			args: args{
				x: 0,
				y: math.MinInt,
			},
			want: want{
				overflow: true,
			},
		},
		{
			name: "min;min",
			args: args{
				x: math.MinInt,
				y: math.MinInt,
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "min;max",
			args: args{
				x: math.MinInt,
				y: math.MaxInt,
			},
			want: want{
				overflow: true,
			},
		},
		{
			name: "max;min",
			args: args{
				x: math.MaxInt,
				y: math.MinInt,
			},
			want: want{
				overflow: true,
			},
		},
		{
			name: "max;max",
			args: args{
				x: math.MaxInt,
				y: math.MaxInt,
			},
			want: want{
				result: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := Sub(tt.args.x, tt.args.y)
			if !tt.want.overflow {
				if !ok {
					t.Fatalf("unexpected overflow")
				}
				if result != tt.want.result {
					t.Fatalf("%d expected, got %d", tt.want.result, result)
				}
			} else {
				if ok {
					t.Fatalf("overflow expected, got %d", result)
				}
			}
		})
	}
}
