// https://medium.com/@dreissenzahn/functional-options-in-go-35279f535c6
package main
import "fmt"
type Foo struct {
	Option0 string
	Option1 bool
	Option2 int
}
type option func(*Foo)
func Option1(option1 bool) option {
	return func(f *Foo) {
		f.Option1 = option1
	}
}
func Option2(option2 int) option {
	return func(f *Foo) {
		f.Option2 = option2
	}
}

// t- we use functional options in the constructor to set the default behaviour
func NewFoo(option0 string, opts ...option) *Foo {
	// Default options
	f := &Foo{
		Option0: option0,
		Option1: false,
		Option2: 10,
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}
func main() {
	f := NewFoo("foo", Option1(true), Option2(20))
	fmt.Println(f.Option0) // Prints "foo"
	fmt.Println(f.Option1) // Prints true
	fmt.Println(f.Option2) // Prints 20
}
