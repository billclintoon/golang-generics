package golang_generics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello" + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		"bill",
		"clinton",
	}
	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		"bill",
		"clinton",
	}

	assert.Equal(t, "bill", data.ChangeFirst("bill"))
	assert.Equal(t, "Hello bill", data.SayHello("bill"))
}
