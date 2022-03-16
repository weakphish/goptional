package goptional

type Option[T any] struct {
	exists bool
	some   T
	none   error
}

// Return a new Option type, filled with the value and an optional error. If an error is provided,
// the value is erased and only the error will be visible. Because of this, you can provide any
// dummy value for T if the Option will be an error.
func New[T any](val T, err error) Option[T] {
	var o Option[T]
	if err != nil {
		var t T
		o.exists = false
		o.some = t
		o.none = err
	} else {
		o.exists = true
		o.some = val
		o.none = nil
	}
	return o
}

// Unwrap the option - if the option has an error, it will be consumed and passed into a `panic`
// call. Otherwise, a value of type T will be returned.
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

// Unwraps the option, but instead of panicking if there is nothing to be done, computes a provided
// function should it fail and returns the result of it.
func (o *Option[T]) UnwrapOrElse(f func() T) T {
	if !o.exists {
		val := f()
		return val
	} else {
		return o.some
	}
}
