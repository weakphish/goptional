package main

import (
	"errors"
	"math/rand"

	"github.com/weakphish/goptional"
)

func generateEvenNumber[T any]() goptional.Option[int] {
	num := rand.Intn(100)
	if num%2 == 0 {
		return goptional.New(num, nil)
	} else {
		return goptional.New(num, errors.New("Not an even number."))
	}
}

func main() {
	o := generateEvenNumber[int]()
	val := o.Unwrap()
	println(val)
}
