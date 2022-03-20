package goptional

type Option[T any] struct {
	exists bool
	some   T
	none   error
}

func None[T any](err error) *Option[T] {
	var o Option[T]
	var t T
	o.exists = false
	// use the default value of type T as a placeholder for the Option's Some, which will not
	// be accessed
	o.some = t
	o.none = err

	return &o
}

func Some[T any](value T) *Option[T] {
	var o Option[T]
	o.exists = true
	o.some = value
	o.none = nil

	return &o
}

// Unwrap the option - if the underlying value is None, the attached error is passed into a `panic`
// call. Otherwise, the underlying Some of type T will be returned.
func (o *Option[T]) Unwrap() T {
	if !o.exists {
		panic(o.none)
	} else {
		return o.some
	}
}

// Returns a provided default value if there is no Some.
func (o *Option[T]) UnwrapOr(def T) T {
	if !o.exists {
		return def
	} else {
		return o.some
	}
}

// Unwraps the option, but instead of panicking if there is None, computes a provided
// function should it fail and returns the result of it.
func (o *Option[T]) UnwrapOrElse(f func() T) T {
	if !o.exists {
		val := f()
		return val
	} else {
		return o.some
	}
}
