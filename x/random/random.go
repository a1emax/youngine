package random

import (
	"github.com/a1emax/youngine/fault"
)

// Random number generator.
type Random interface {

	// Seed uses the provided seed value to initialize generator to deterministic state.
	//
	// Seed should not be called concurrently with any other [Random] method.
	Seed(seed int64)

	// Int returns non-negative pseudo-random integer number.
	Int() int

	// Intn returns non-negative pseudo-random integer number in half-open interval [0,n).
	//
	// Intn panics if n <= 0.
	Intn(n int) int

	// Int31 returns non-negative pseudo-random 31-bit integer number.
	Int31() int32

	// Int31n returns non-negative pseudo-random 31-bit integer number in half-open interval [0,n).
	//
	// Int31n panics if n <= 0.
	Int31n(n int32) int32

	// Int63 returns non-negative pseudo-random 63-bit integer number.
	Int63() int64

	// Int63n returns non-negative pseudo-random 63-bit integer number in half-open interval [0,n).
	//
	// Int63n panics if n <= 0.
	Int63n(n int64) int64

	// Uint32 returns pseudo-random 32-bit integer number.
	Uint32() uint32

	// Uint64 returns pseudo-random 64-bit integer number.
	Uint64() uint64

	// Float32 returns pseudo-random 32-bit floating-point number in half-open interval [0.0,1.0).
	Float32() float32

	// Float64 returns pseudo-random 64-bit floating-point number in half-open interval [0.0,1.0).
	Float64() float64

	// NormFloat64 returns normally distributed 64-bit floating-point number in range -[math.MaxFloat64]
	// through +[math.MaxFloat64] inclusive, with standard normal distribution (mean = 0, stddev = 1).
	// To produce a different normal distribution, callers can adjust the output using:
	//
	//	sample = NormFloat64() * desiredStdDev + desiredMean
	NormFloat64() float64

	// ExpFloat64 returns exponentially distributed 64-bit floating-point number in range (0, +[math.MaxFloat64]]
	// with an exponential distribution whose rate parameter (lambda) is 1 and whose mean is 1/lambda (1).
	// To produce a distribution with a different rate parameter, callers can adjust the output using:
	//
	//	sample = ExpFloat64() / desiredRateParameter
	ExpFloat64() float64

	// Perm returns pseudo-random permutation of integer numbers in the half-open interval [0,n).
	//
	// Perm panics if n < 0.
	Perm(n int) []int

	// Shuffle pseudo-randomizes the order of n elements using given function to swap elements with indexes i and j.
	//
	// Shuffle panics if n < 0.
	Shuffle(n int, swap func(i, j int))

	// Read generates len(p) random bytes and writes them into p. It always returns len(p) and a nil error.
	//
	// Read should not be called concurrently with any other [Random] method.
	Read(p []byte) (n int, err error)
}

// Intw returns non-negative pseudo-random integer number in half-open interval [0,len(w)) considering weights w.
//
// Intw treats negative weights as zero ones.
func Intw(r Random, w ...float64) int {
	if r == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	var sum float64
	for _, v := range w {
		sum += max(0, v)
	}

	x := r.Float64() * sum
	for i, v := range w {
		if x < v {
			return i
		}

		x -= v
	}

	return 0 // should never happen
}

// Prob returns true with probability p.
//
// Prob always returns false if p <= 0, and always returns true if p >= 1.
func Prob(r Random, p float64) bool {
	return r.Float64() < p
}
