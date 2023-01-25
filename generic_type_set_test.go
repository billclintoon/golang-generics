package golang_generics

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestBagString(t *testing.T) {
	name := Bag[string]{"bill"}
	PrintBag(name)
}

func TestBagInt(t *testing.T) {
	numbers := Bag[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	PrintBag(numbers)
}
