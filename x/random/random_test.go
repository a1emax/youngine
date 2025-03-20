package random

import (
	"testing"
)

const testEpsilon = 1e-10

type testRandom struct {
	_Seed        func(seed int64)
	_Int         func() int
	_Intn        func(n int) int
	_Int31       func() int32
	_Int31n      func(n int32) int32
	_Int63       func() int64
	_Int63n      func(n int64) int64
	_Uint32      func() uint32
	_Uint64      func() uint64
	_Float32     func() float32
	_Float64     func() float64
	_NormFloat64 func() float64
	_ExpFloat64  func() float64
	_Perm        func(n int) []int
	_Shuffle     func(n int, swap func(i, j int))
	_Read        func(p []byte) (n int, err error)
}

func (r testRandom) Seed(seed int64) {
	r._Seed(seed)
}

func (r testRandom) Int() int {
	return r._Int()
}

func (r testRandom) Intn(n int) int {
	return r._Intn(n)
}

func (r testRandom) Int31() int32 {
	return r._Int31()
}

func (r testRandom) Int31n(n int32) int32 {
	return r._Int31n(n)
}

func (r testRandom) Int63() int64 {
	return r._Int63()
}

func (r testRandom) Int63n(n int64) int64 {
	return r._Int63n(n)
}

func (r testRandom) Uint32() uint32 {
	return r._Uint32()
}

func (r testRandom) Uint64() uint64 {
	return r._Uint64()
}

func (r testRandom) Float32() float32 {
	return r._Float32()
}

func (r testRandom) Float64() float64 {
	return r._Float64()
}

func (r testRandom) NormFloat64() float64 {
	return r._NormFloat64()
}

func (r testRandom) ExpFloat64() float64 {
	return r._ExpFloat64()
}

func (r testRandom) Perm(n int) []int {
	return r._Perm(n)
}

func (r testRandom) Shuffle(n int, swap func(i, j int)) {
	r._Shuffle(n, swap)
}

func (r testRandom) Read(p []byte) (n int, err error) {
	return r._Read(p)
}

func TestIntw(t *testing.T) {
	type args struct {
		r Random
		w []float64
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
			name: "<30+0%;30,0,25,45",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.30 + 0 - testEpsilon
					},
				},
				w: []float64{30, 0, 25, 45},
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "30+0%;30,0,25,40",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.30 + 0
					},
				},
				w: []float64{30, 0, 25, 45},
			},
			want: want{
				result: 2,
			},
		},
		{
			name: "30+0+25%;30,0,25,45",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.30 + 0 + 0.25
					},
				},
				w: []float64{30, 0, 25, 45},
			},
			want: want{
				result: 3,
			},
		},
		{
			name: "0%;10,20,30,40",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0
					},
				},
				w: []float64{10, 20, 30, 40},
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "<10%;10,20,30,40",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.10 - testEpsilon
					},
				},
				w: []float64{10, 20, 30, 40},
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "10%;10,20,30,40",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.10
					},
				},
				w: []float64{10, 20, 30, 40},
			},
			want: want{
				result: 1,
			},
		},
		{
			name: "<10+20%;10,20,30,40",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.10 + 0.20 - testEpsilon
					},
				},
				w: []float64{10, 20, 30, 40},
			},
			want: want{
				result: 1,
			},
		},
		{
			name: "10+20%;10,20,30,40",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.10 + 0.20
					},
				},
				w: []float64{10, 20, 30, 40},
			},
			want: want{
				result: 2,
			},
		},
		{
			name: "<10+20+30%;10,20,30,40",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.10 + 0.20 + 0.30 - testEpsilon
					},
				},
				w: []float64{10, 20, 30, 40},
			},
			want: want{
				result: 2,
			},
		},
		{
			name: "10+20+30%;10,20,30,40",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.10 + 0.20 + 0.30
					},
				},
				w: []float64{10, 20, 30, 40},
			},
			want: want{
				result: 3,
			},
		},
		{
			name: "<100%;10,20,30,40",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 1 - testEpsilon
					},
				},
				w: []float64{10, 20, 30, 40},
			},
			want: want{
				result: 3,
			},
		},
		{
			name: "0%;40,30,20,10",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0
					},
				},
				w: []float64{40, 30, 20, 10},
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "<40%;40,30,20,10",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.40 - testEpsilon
					},
				},
				w: []float64{40, 30, 20, 10},
			},
			want: want{
				result: 0,
			},
		},
		{
			name: "40%;40,30,20,10",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.40
					},
				},
				w: []float64{40, 30, 20, 10},
			},
			want: want{
				result: 1,
			},
		},
		{
			name: "<40+30%;40,30,20,10",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.40 + 0.30 - testEpsilon
					},
				},
				w: []float64{40, 30, 20, 10},
			},
			want: want{
				result: 1,
			},
		},
		{
			name: "40+30%;40,30,20,10",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.40 + 0.30
					},
				},
				w: []float64{40, 30, 20, 10},
			},
			want: want{
				result: 2,
			},
		},
		{
			name: "<40+30+20%;40,30,20,10",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.40 + 0.30 + 0.20 - testEpsilon
					},
				},
				w: []float64{40, 30, 20, 10},
			},
			want: want{
				result: 2,
			},
		},
		{
			name: "40+30+20%;40,30,20,10",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.40 + 0.30 + 0.20
					},
				},
				w: []float64{40, 30, 20, 10},
			},
			want: want{
				result: 3,
			},
		},
		{
			name: "<100%;40,30,20,10",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 1 - testEpsilon
					},
				},
				w: []float64{40, 30, 20, 10},
			},
			want: want{
				result: 3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Intw(tt.args.r, tt.args.w...)
			if result != tt.want.result {
				t.Fatalf("%d expected, got %d", tt.want.result, result)
			}
		})
	}
}

func TestProb(t *testing.T) {
	type args struct {
		r Random
		p float64
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
			name: "0;<0",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0
					},
				},
				p: 0 - testEpsilon,
			},
			want: want{
				result: false,
			},
		},
		{
			name: "0.3;<0",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.3
					},
				},
				p: 0 - testEpsilon,
			},
			want: want{
				result: false,
			},
		},
		{
			name: "<1;<0",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 1 - testEpsilon
					},
				},
				p: 0 - testEpsilon,
			},
			want: want{
				result: false,
			},
		},
		{
			name: "0;0",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0
					},
				},
				p: 0,
			},
			want: want{
				result: false,
			},
		},
		{
			name: "0.3;0",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.3
					},
				},
				p: 0,
			},
			want: want{
				result: false,
			},
		},
		{
			name: "<1;0",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 1 - testEpsilon
					},
				},
				p: 0,
			},
			want: want{
				result: false,
			},
		},
		{
			name: "<0.3;0.3",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.3 - testEpsilon
					},
				},
				p: 0.3,
			},
			want: want{
				result: true,
			},
		},
		{
			name: "0.3;0.3",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.3
					},
				},
				p: 0.3,
			},
			want: want{
				result: false,
			},
		},
		{
			name: "<1;0.3",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 1 - testEpsilon
					},
				},
				p: 0.3,
			},
			want: want{
				result: false,
			},
		},
		{
			name: "<0.3;1",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.3 - testEpsilon
					},
				},
				p: 1,
			},
			want: want{
				result: true,
			},
		},
		{
			name: "0.3;1",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.3
					},
				},
				p: 1,
			},
			want: want{
				result: true,
			},
		},
		{
			name: "<1;1",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 1 - testEpsilon
					},
				},
				p: 1,
			},
			want: want{
				result: true,
			},
		},
		{
			name: "0;1",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0
					},
				},
				p: 1,
			},
			want: want{
				result: true,
			},
		},
		{
			name: "0.3;1",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.3
					},
				},
				p: 1,
			},
			want: want{
				result: true,
			},
		},
		{
			name: "<1;1",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 1 - testEpsilon
					},
				},
				p: 1,
			},
			want: want{
				result: true,
			},
		},
		{
			name: "0;>1",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0
					},
				},
				p: 1 + testEpsilon,
			},
			want: want{
				result: true,
			},
		},
		{
			name: "0.3;>1",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 0.3
					},
				},
				p: 1 + testEpsilon,
			},
			want: want{
				result: true,
			},
		},
		{
			name: "<1;>1",
			args: args{
				r: testRandom{
					_Float64: func() float64 {
						return 1 - testEpsilon
					},
				},
				p: 1 + testEpsilon,
			},
			want: want{
				result: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Prob(tt.args.r, tt.args.p)
			if result != tt.want.result {
				t.Fatalf("%t expected, got %t", tt.want.result, result)
			}
		})
	}
}
