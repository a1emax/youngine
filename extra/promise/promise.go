package promise

import (
	"errors"

	"github.com/a1emax/youngine/fault"
)

// Promise represents eventual completion (or failure) of asynchronous operation and its resulting value of type T.
type Promise[T any] interface {
	Thener

	// Value returns resulting value of asynchronous operation.
	//
	// Value panics if operation is still in progress or failed.
	Value() T
}

// ResolveFunc notifies that asynchronous operation was completed successfully with given resulting value.
//
// ResolveFunc panics if called more than once or after corresponding [RejectFunc] has already been called.
type ResolveFunc[T any] func(value T)

// RejectFunc notifies that  asynchronous operation failed by given reason (if reason is nil, ErrRejected
// will be used instead).
//
// RejectFunc panics if called more than once or after corresponding [ResolveFunc] has already been called.
type RejectFunc func(err error)

// promiseImpl is the implementation of the [Promise] interface.
type promiseImpl[T any] struct {
	status    status
	reactions []promiseReaction
	value     T
	err       error
}

// status of asynchronous operation.
type status int

const (

	// statusPending indicates that asynchronous operation is still in progress.
	statusPending status = iota

	// statusResolved indicates that asynchronous operation was completed successfully.
	statusResolved

	// statusRejected indicates that asynchronous operation failed.
	statusRejected
)

// promiseReaction is enqueued pair of handlers. See [Thener] for details.
type promiseReaction struct {
	onResolved OnResolvedFunc
	onRejected OnRejectedFunc
}

// New initializes and returns new [Promise] along with new functions that settle it.
func New[T any]() (Promise[T], ResolveFunc[T], RejectFunc) {
	p := &promiseImpl[T]{}

	resolve := func(value T) {
		if p == nil {
			panic(fault.Trace(fault.ErrInvalidUse))
		}
		defer func() {
			p = nil
		}()

		p.status = statusResolved
		p.value = value
		p.processReactions()
	}

	reject := func(err error) {
		if p == nil {
			panic(fault.Trace(fault.ErrInvalidUse))
		}
		defer func() {
			p = nil
		}()

		if err == nil {
			err = ErrRejected
		}

		p.status = statusRejected
		p.err = err
		p.processReactions()
	}

	return p, resolve, reject
}

// Then implements the [Thener] interface.
func (p *promiseImpl[T]) Then(onResolved OnResolvedFunc, onRejected OnRejectedFunc) {
	if onResolved == nil && onRejected == nil {
		return
	}

	if p.status == statusPending {
		p.reactions = append(p.reactions, promiseReaction{onResolved, onRejected})
	} else {
		p.callHandler(onResolved, onRejected)
	}
}

// Value implements the [Promise] interface.
func (p *promiseImpl[T]) Value() T {
	if p.status != statusResolved {
		panic(fault.Trace(fault.ErrInvalidUse))
	}

	return p.value
}

// processReactions, after asynchronous operation is finished, for each pair of enqueued handlers calls one of them,
// depending on final status of operation.
//
// processReactions panics after calling all relevant handlers if any of them panicked.
func (p *promiseImpl[T]) processReactions() {
	defer func() {
		p.reactions = nil
	}()

	var err error
	for _, reaction := range p.reactions {
		err = errors.Join(err, fault.Recover(func() {
			p.callHandler(reaction.onResolved, reaction.onRejected)
		}))
	}
	if err != nil {
		panic(err)
	}
}

// callHandler, after asynchronous operation is finished, calls one of given handlers,
// depending on final status of operation.
func (p *promiseImpl[T]) callHandler(onResolved OnResolvedFunc, onRejected OnRejectedFunc) {
	switch p.status {
	case statusResolved:
		if onResolved != nil {
			onResolved()
		}
	case statusRejected:
		if onRejected != nil {
			onRejected(p.err)
		}
	}
}
