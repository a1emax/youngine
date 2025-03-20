package basic

import (
	"image"
	"testing"
)

func TestRectBtw(t *testing.T) {
	type args struct {
		p0 Vec2
		p1 Vec2
	}
	type want struct {
		result Rect
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(0,0);(0,0)",
			args: args{
				p0: Vec2{0, 0},
				p1: Vec2{0, 0},
			},
			want: want{
				result: Rect{Vec2{0, 0}, Vec2{0, 0}},
			},
		},
		{
			name: "(1,20);(10,2)",
			args: args{
				p0: Vec2{1, 20},
				p1: Vec2{10, 2},
			},
			want: want{
				result: Rect{Vec2{1, 2}, Vec2{9, 18}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RectBtw(tt.args.p0, tt.args.p1)
			if result != tt.want.result {
				t.Fatalf("%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestRect_Image(t *testing.T) {
	type args struct {
		r Rect
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
				r: Rect{Vec2{0, 0}, Vec2{0, 0}},
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
				r: Rect{Vec2{1, 2}, Vec2{10, 20}},
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

func TestRect_IsEmpty(t *testing.T) {
	type args struct {
		r Rect
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
				r: Rect{},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(0,0)+(0,0)",
			args: args{
				r: Rect{Vec2{0, 0}, Vec2{0, 0}},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(0,0)+(-1,-1)",
			args: args{
				r: Rect{Vec2{0, 0}, Vec2{-1, -1}},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(0,0)+(1,1)",
			args: args{
				r: Rect{Vec2{0, 0}, Vec2{1, 1}},
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

func TestRect_Left(t *testing.T) {
	type args struct {
		r Rect
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
			name: "(1,2)+(10,20)",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{10, 20}},
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
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestRect_Top(t *testing.T) {
	type args struct {
		r Rect
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
			name: "(1,2)+(10,20)",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{10, 20}},
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
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestRect_Right(t *testing.T) {
	type args struct {
		r Rect
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
			name: "(1,2)+(10,20)",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{10, 20}},
			},
			want: want{
				result: 11,
			},
		},
		{
			name: "(1,2)+(-10,-20)",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{-10, -20}},
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
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestRect_Bottom(t *testing.T) {
	type args struct {
		r Rect
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
			name: "(1,2)+(10,20)",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{10, 20}},
			},
			want: want{
				result: 22,
			},
		},
		{
			name: "(1,2)+(-10,-20)",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{-10, -20}},
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
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestRect_Width(t *testing.T) {
	type args struct {
		r Rect
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
			name: "(1,2)+(10,20)",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{10, 20}},
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
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestRect_Height(t *testing.T) {
	type args struct {
		r Rect
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
			name: "(1,2)+(10,20)",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{10, 20}},
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
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}

func TestRect_Contains(t *testing.T) {
	type args struct {
		r Rect
		p Vec2
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
				r: Rect{Vec2{0, 0}, Vec2{0, 0}},
				p: Vec2{0, 0},
			},
			want: want{
				result: false,
			},
		},
		{
			name: "(0,0)+(1,1);(0,0)",
			args: args{
				r: Rect{Vec2{0, 0}, Vec2{1, 1}},
				p: Vec2{0, 0},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(0,0)+(1,1);(0.9,0.9)",
			args: args{
				r: Rect{Vec2{0, 0}, Vec2{1, 1}},
				p: Vec2{0.9, 0.9},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(0,0)+(1,1);(1,1)",
			args: args{
				r: Rect{Vec2{0, 0}, Vec2{1, 1}},
				p: Vec2{1, 1},
			},
			want: want{
				result: false,
			},
		},
		{
			name: "(-1,-1)+(2,2);(0,0)",
			args: args{
				r: Rect{Vec2{-1, -1}, Vec2{2, 2}},
				p: Vec2{0, 0},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(-5,-5)+(1,1);(-5,-5)",
			args: args{
				r: Rect{Vec2{-5, -5}, Vec2{1, 1}},
				p: Vec2{-5, -5},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(-5,-5)+(1,1);(-4.1,-4.1)",
			args: args{
				r: Rect{Vec2{-5, -5}, Vec2{1, 1}},
				p: Vec2{-4.1, -4.1},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(-5,-5)+(1,1);(-4,-4)",
			args: args{
				r: Rect{Vec2{-5, -5}, Vec2{1, 1}},
				p: Vec2{-4, -4},
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

func TestRect_Inner(t *testing.T) {
	type args struct {
		r Rect
	}
	type want struct {
		result Rect
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(0,0)+(0,0)",
			args: args{
				r: Rect{Vec2{0, 0}, Vec2{0, 0}},
			},
			want: want{
				result: Rect{Vec2{0, 0}, Vec2{0, 0}},
			},
		},
		{
			name: "(1,2)+(10,20)",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{10, 20}},
			},
			want: want{
				result: Rect{Vec2{0, 0}, Vec2{10, 20}},
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

func TestRect_Move(t *testing.T) {
	type args struct {
		r Rect
		d Vec2
	}
	type want struct {
		result Rect
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,2)+(10,20);(0,0)",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{10, 20}},
				d: Vec2{0, 0},
			},
			want: want{
				result: Rect{Vec2{1, 2}, Vec2{10, 20}},
			},
		},
		{
			name: "(1,2)+(10,20);(3,5)",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{10, 20}},
				d: Vec2{3, 5},
			},
			want: want{
				result: Rect{Vec2{4, 7}, Vec2{10, 20}},
			},
		},
		{
			name: "(1,2)+(10,20);(-3,-5)",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{10, 20}},
				d: Vec2{-3, -5},
			},
			want: want{
				result: Rect{Vec2{-2, -3}, Vec2{10, 20}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.r.Move(tt.args.d)
			if result != tt.want.result {
				t.Fatalf("(%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestRect_MoveNum(t *testing.T) {
	type args struct {
		r Rect
		d Float
	}
	type want struct {
		result Rect
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,2)+(10,20);0",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{10, 20}},
				d: 0,
			},
			want: want{
				result: Rect{Vec2{1, 2}, Vec2{10, 20}},
			},
		},
		{
			name: "(1,2)+(10,20);3",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{10, 20}},
				d: 3,
			},
			want: want{
				result: Rect{Vec2{4, 5}, Vec2{10, 20}},
			},
		},
		{
			name: "(1,2)+(10,20);-3",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{10, 20}},
				d: -3,
			},
			want: want{
				result: Rect{Vec2{-2, -1}, Vec2{10, 20}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.r.MoveNum(tt.args.d)
			if result != tt.want.result {
				t.Fatalf("(%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestRect_Overlaps(t *testing.T) {
	type args struct {
		r Rect
		s Rect
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
				r: Rect{Vec2{0, 0}, Vec2{0, 0}},
				s: Rect{Vec2{0, 0}, Vec2{0, 0}},
			},
			want: want{
				result: false,
			},
		},
		{
			name: "(0,0)+(1,1);(0,0)+(0,0)",
			args: args{
				r: Rect{Vec2{0, 0}, Vec2{1, 1}},
				s: Rect{Vec2{0, 0}, Vec2{0, 0}},
			},
			want: want{
				result: false,
			},
		},
		{
			name: "(0,0)+(0,0);(0,0)+(1,1)",
			args: args{
				r: Rect{Vec2{0, 0}, Vec2{0, 0}},
				s: Rect{Vec2{0, 0}, Vec2{1, 1}},
			},
			want: want{
				result: false,
			},
		},
		{
			name: "(0,0)+(1,1);(0,0)+(1,1)",
			args: args{
				r: Rect{Vec2{0, 0}, Vec2{1, 1}},
				s: Rect{Vec2{0, 0}, Vec2{1, 1}},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(0,0)+(1,1);(1,1)+(1,1)",
			args: args{
				r: Rect{Vec2{0, 0}, Vec2{1, 1}},
				s: Rect{Vec2{1, 1}, Vec2{1, 1}},
			},
			want: want{
				result: false,
			},
		},
		{
			name: "(-5,2)+(5,3);(-4,3)+(3,2)",
			args: args{
				r: Rect{Vec2{-5, 2}, Vec2{5, 3}},
				s: Rect{Vec2{-4, 3}, Vec2{3, 2}},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(-5,2)+(5,3);(-9,0)+(5,3)",
			args: args{
				r: Rect{Vec2{-5, 2}, Vec2{5, 3}},
				s: Rect{Vec2{-9, 0}, Vec2{5, 3}},
			},
			want: want{
				result: true,
			},
		},
		{
			name: "(-5,2)+(5,3);(-9,0)+(2,1)",
			args: args{
				r: Rect{Vec2{-5, 2}, Vec2{5, 3}},
				s: Rect{Vec2{-9, 0}, Vec2{2, 1}},
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

func TestRect_Resize(t *testing.T) {
	type args struct {
		r Rect
		d Vec2
	}
	type want struct {
		result Rect
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,2)+(10,20);(0,0)",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{10, 20}},
				d: Vec2{0, 0},
			},
			want: want{
				result: Rect{Vec2{1, 2}, Vec2{10, 20}},
			},
		},
		{
			name: "(1,2)+(10,20);(3,5)",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{10, 20}},
				d: Vec2{3, 5},
			},
			want: want{
				result: Rect{Vec2{-2, -3}, Vec2{16, 30}},
			},
		},
		{
			name: "(1,2)+(10,20);(-3,-5)",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{10, 20}},
				d: Vec2{-3, -5},
			},
			want: want{
				result: Rect{Vec2{4, 7}, Vec2{4, 10}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.r.Resize(tt.args.d)
			if result != tt.want.result {
				t.Fatalf("(%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestRect_ResizeNum(t *testing.T) {
	type args struct {
		r Rect
		d Float
	}
	type want struct {
		result Rect
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "(1,2)+(10,20);0",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{10, 20}},
				d: 0,
			},
			want: want{
				result: Rect{Vec2{1, 2}, Vec2{10, 20}},
			},
		},
		{
			name: "(1,2)+(10,20);3",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{10, 20}},
				d: 3,
			},
			want: want{
				result: Rect{Vec2{-2, -1}, Vec2{16, 26}},
			},
		},
		{
			name: "(1,2)+(10,20);-3",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{10, 20}},
				d: -3,
			},
			want: want{
				result: Rect{Vec2{4, 5}, Vec2{4, 14}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.r.ResizeNum(tt.args.d)
			if result != tt.want.result {
				t.Fatalf("(%s expected, got %s", tt.want.result, result)
			}
		})
	}
}

func TestRect_Square(t *testing.T) {
	type args struct {
		r Rect
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
			name: "(1,2)+(10,20)",
			args: args{
				r: Rect{Vec2{1, 2}, Vec2{10, 20}},
			},
			want: want{
				result: 200,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.args.r.Square()
			if result != tt.want.result {
				t.Fatalf("(%g expected, got %g", tt.want.result, result)
			}
		})
	}
}
