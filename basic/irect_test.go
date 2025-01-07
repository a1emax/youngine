package basic

import (
	"image"
	"testing"
)

func TestIrect_Image(t *testing.T) {
	type args struct {
		r Irect
	}
	type want struct {
		result image.Rectangle
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(0,0)+(0,0)",
			args: args{
				r: Irect{Ivec2{0, 0}, Ivec2{0, 0}},
			},
			want: want{
				result: image.Rectangle{
					Min: image.Point{X: 0, Y: 0},
					Max: image.Point{X: 0, Y: 0},
				},
			},
		},
		{
			name: "(1,2)+(10,20)",
			args: args{
				r: Irect{Ivec2{1, 2}, Ivec2{10, 20}},
			},
			want: want{
				result: image.Rectangle{
					Min: image.Point{X: 1, Y: 2},
					Max: image.Point{X: 11, Y: 22},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.r.Image()
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestIrect_IsEmpty(t *testing.T) {
	type args struct {
		r Irect
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
				r: Irect{},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(0,0)+(0,0)",
			args: args{
				r: Irect{Ivec2{0, 0}, Ivec2{0, 0}},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(0,0)+(-1,-1)",
			args: args{
				r: Irect{Ivec2{0, 0}, Ivec2{-1, -1}},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(0,0)+(1,1)",
			args: args{
				r: Irect{Ivec2{0, 0}, Ivec2{1, 1}},
			},
			want: want{
				result: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.r.IsEmpty()
			if result != tt.want.result {
				t.Fatalf("(%t expected, got %t", tt.want.result, result)
			}
		})
	}
}

func TestIrect_Left(t *testing.T) {
	type args struct {
		r Irect
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
			name: "(1,2)+(10,20)",
			args: args{
				r: Irect{Ivec2{1, 2}, Ivec2{10, 20}},
			},
			want: want{
				result: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.r.Left()
			if result != tt.want.result {
				t.Fatalf("(%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestIrect_Top(t *testing.T) {
	type args struct {
		r Irect
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
			name: "(1,2)+(10,20)",
			args: args{
				r: Irect{Ivec2{1, 2}, Ivec2{10, 20}},
			},
			want: want{
				result: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.r.Top()
			if result != tt.want.result {
				t.Fatalf("(%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestIrect_Right(t *testing.T) {
	type args struct {
		r Irect
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
			name: "(1,2)+(10,20)",
			args: args{
				r: Irect{Ivec2{1, 2}, Ivec2{10, 20}},
			},
			want: want{
				result: 11,
			},
		},
		{
			name: "(1,2)+(-10,-20)",
			args: args{
				r: Irect{Ivec2{1, 2}, Ivec2{-10, -20}},
			},
			want: want{
				result: -9,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.r.Right()
			if result != tt.want.result {
				t.Fatalf("(%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestIrect_Bottom(t *testing.T) {
	type args struct {
		r Irect
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
			name: "(1,2)+(10,20)",
			args: args{
				r: Irect{Ivec2{1, 2}, Ivec2{10, 20}},
			},
			want: want{
				result: 22,
			},
		},
		{
			name: "(1,2)+(-10,-20)",
			args: args{
				r: Irect{Ivec2{1, 2}, Ivec2{-10, -20}},
			},
			want: want{
				result: -18,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.r.Bottom()
			if result != tt.want.result {
				t.Fatalf("(%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestIrect_Width(t *testing.T) {
	type args struct {
		r Irect
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
			name: "(1,2)+(10,20)",
			args: args{
				r: Irect{Ivec2{1, 2}, Ivec2{10, 20}},
			},
			want: want{
				result: 10,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.r.Width()
			if result != tt.want.result {
				t.Fatalf("(%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestIrect_Height(t *testing.T) {
	type args struct {
		r Irect
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
			name: "(1,2)+(10,20)",
			args: args{
				r: Irect{Ivec2{1, 2}, Ivec2{10, 20}},
			},
			want: want{
				result: 20,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.r.Height()
			if result != tt.want.result {
				t.Fatalf("(%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestIrect_Inner(t *testing.T) {
	type args struct {
		r Irect
	}
	type want struct {
		result Irect
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(0,0)+(0,0)",
			args: args{
				r: Irect{Ivec2{0, 0}, Ivec2{0, 0}},
			},
			want: want{
				result: Irect{Ivec2{0, 0}, Ivec2{0, 0}},
			},
		},
		{
			name: "(1,2)+(10,20)",
			args: args{
				r: Irect{Ivec2{1, 2}, Ivec2{10, 20}},
			},
			want: want{
				result: Irect{Ivec2{0, 0}, Ivec2{10, 20}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.r.Inner()
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestIrect_Contains(t *testing.T) {
	type args struct {
		r Irect
		p Ivec2
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
			name: "(0,0)+(0,0);(0,0)",
			args: args{
				r: Irect{Ivec2{0, 0}, Ivec2{0, 0}},
				p: Ivec2{0, 0},
			},
			want: want{
				result: false,
			},
		},
		{
			name: "(0,0)+(1,1);(0,0)",
			args: args{
				r: Irect{Ivec2{0, 0}, Ivec2{1, 1}},
				p: Ivec2{0, 0},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(0,0)+(1,1);(1,1)",
			args: args{
				r: Irect{Ivec2{0, 0}, Ivec2{1, 1}},
				p: Ivec2{1, 1},
			},
			want: want{
				result: false,
			},
		},
		{
			name: "(-1,-1)+(2,2);(0,0)",
			args: args{
				r: Irect{Ivec2{-1, -1}, Ivec2{2, 2}},
				p: Ivec2{0, 0},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(-5,-5)+(1,1);(-5,-5)",
			args: args{
				r: Irect{Ivec2{-5, -5}, Ivec2{1, 1}},
				p: Ivec2{-5, -5},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(-5,-5)+(1,1);(-4,-4)",
			args: args{
				r: Irect{Ivec2{-5, -5}, Ivec2{1, 1}},
				p: Ivec2{-4, -4},
			},
			want: want{
				result: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.r.Contains(tt.args.p)
			if result != tt.want.result {
				t.Fatalf("(%t expected, got %t", tt.want.result, result)
			}
		})
	}
}

func TestIrect_Overlaps(t *testing.T) {
	type args struct {
		r Irect
		s Irect
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
			name: "(0,0)+(0,0);(0,0)+(0,0)",
			args: args{
				r: Irect{Ivec2{0, 0}, Ivec2{0, 0}},
				s: Irect{Ivec2{0, 0}, Ivec2{0, 0}},
			},
			want: want{
				result: false,
			},
		},
		{
			name: "(0,0)+(1,1);(0,0)+(0,0)",
			args: args{
				r: Irect{Ivec2{0, 0}, Ivec2{1, 1}},
				s: Irect{Ivec2{0, 0}, Ivec2{0, 0}},
			},
			want: want{
				result: false,
			},
		},
		{
			name: "(0,0)+(0,0);(0,0)+(1,1)",
			args: args{
				r: Irect{Ivec2{0, 0}, Ivec2{0, 0}},
				s: Irect{Ivec2{0, 0}, Ivec2{1, 1}},
			},
			want: want{
				result: false,
			},
		},
		{
			name: "(0,0)+(1,1);(0,0)+(1,1)",
			args: args{
				r: Irect{Ivec2{0, 0}, Ivec2{1, 1}},
				s: Irect{Ivec2{0, 0}, Ivec2{1, 1}},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(0,0)+(1,1);(1,1)+(1,1)",
			args: args{
				r: Irect{Ivec2{0, 0}, Ivec2{1, 1}},
				s: Irect{Ivec2{1, 1}, Ivec2{1, 1}},
			},
			want: want{
				result: false,
			},
		},
		{
			name: "(-5,2)+(5,3);(-4,3)+(3,2)",
			args: args{
				r: Irect{Ivec2{-5, 2}, Ivec2{5, 3}},
				s: Irect{Ivec2{-4, 3}, Ivec2{3, 2}},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(-5,2)+(5,3);(-9,0)+(5,3)",
			args: args{
				r: Irect{Ivec2{-5, 2}, Ivec2{5, 3}},
				s: Irect{Ivec2{-9, 0}, Ivec2{5, 3}},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(-5,2)+(5,3);(-9,0)+(2,1)",
			args: args{
				r: Irect{Ivec2{-5, 2}, Ivec2{5, 3}},
				s: Irect{Ivec2{-9, 0}, Ivec2{2, 1}},
			},
			want: want{
				result: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.r.Overlaps(tt.args.s)
			if result != tt.want.result {
				t.Fatalf("(%t expected, got %t", tt.want.result, result)
			}
		})
	}
}
