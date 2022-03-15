package goptional

type Option[T any] struct {
	exists bool
	some   T
	none   error
}

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

func (o *Option[T]) Unwrap() T {
	if !o.exists {
		panic(o.none)
	} else {
		return o.some
	}
}
