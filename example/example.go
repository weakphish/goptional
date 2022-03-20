package main

import (
	"errors"

	"github.com/weakphish/goptional"
)

func divide[T any](numerator, denominator int) *goptional.Option[int] {
	if denominator == 0 {
		return goptional.None[int](errors.New("Cannot divide by zero"))
	} else {
		return goptional.Some(numerator / denominator)
	}
}

func main() {
	s := divide[int](6, 3).Unwrap()
	println(s)

	n := divide[int](6, 0).Unwrap()
	println(n)
}
