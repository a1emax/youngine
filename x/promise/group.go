package promise

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
)

// All initializes and returns new promise that will be resolved when all given theners are resolved,
// or will be rejected when any of theners is rejected.
func All(theners ...Thener) Promise[basic.None] {
	for _, thener := range theners {
		if thener == nil {
			panic(fault.Trace(fault.ErrNilPointer))
		}
	}

	p, resolve, reject := New[basic.None]()

	counter := len(theners)
	if counter == 0 {
		resolve(basic.None{})

		return p
	}

	for _, thener := range theners {
		thener.Then(func() {
			if counter == 0 {
				return
			}

			if counter--; counter == 0 {
				resolve(basic.None{})
			}
		}, func(err error) {
			if counter == 0 {
				return
			}

			counter = 0
			reject(ErrRejected)
		})
	}

	return p
}

// Any initializes and returns new promise of index of one of given theners that will be resolved when
// any of theners is resolved, or will be rejected when all theners are rejected.
func Any(theners ...Thener) Promise[int] {
	for _, thener := range theners {
		if thener == nil {
			panic(fault.Trace(fault.ErrNilPointer))
		}
	}

	p, resolve, reject := New[int]()

	counter := len(theners)
	if counter == 0 {
		reject(ErrRejected)

		return p
	}

	for i, thener := range theners {
		i := i

		thener.Then(func() {
			if counter == 0 {
				return
			}

			counter = 0
			resolve(i)
		}, func(err error) {
			if counter == 0 {
				return
			}

			if counter--; counter == 0 {
				reject(ErrRejected)
			}
		})
	}

	return p
}
