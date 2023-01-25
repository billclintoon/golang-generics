package golang_generics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	var result string = length[string]("bill")
	fmt.Println(result)
	var resultNumber int = length[int](100)
	assert.Equal(t, 100, resultNumber)
}
