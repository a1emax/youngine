package promise

// Thener represents eventual completion (or failure) of asynchronous operation.
type Thener interface {

	// Then enqueues pair of handlers to call one of them when asynchronous operation will be finished, or calls
	// one of handlers in place if operation was already finished. Any of handlers can be nil to skip handling
	// corresponding case.
	Then(onResolved OnResolvedFunc, onRejected OnRejectedFunc)
}

// OnResolvedFunc handles case when asynchronous operation was completed successfully.
type OnResolvedFunc func()

// OnRejectedFunc handles case when asynchronous operation failed for given reason.
type OnRejectedFunc func(err error)
