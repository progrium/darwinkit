package objc

import (
	"testing"

	"github.com/progrium/macdriver/internal/assert"
)

// uses NSObject as super without struct tag, regardless of embedded type.
type FooObject struct {
	Object
}

func (f FooObject) MethodForFoo() string {
	return "foo"
}

// uses FooObject from struct tag as super, which should be registered first.
type BarObject struct {
	FooObject `objc:"FooObject"`
}

func (b BarObject) MethodForBar() string {
	return "bar"
}

func TestNewClass(t *testing.T) {
	FooObjectClass := NewClass[FooObject](
		Sel("methodForFoo"),
	)
	RegisterClass(FooObjectClass)

	BarObjectClass := NewClass[BarObject](
		Sel("methodForBar"),
	)
	RegisterClass(BarObjectClass)

	// as per our rules, New returned objects are autoreleased
	WithAutoreleasePool(func() {
		foo := FooObjectClass.New()
		// foo go method
		if foo.MethodForFoo() != "foo" {
			t.Fatal("unexpected return")
		}
		// foo objc method
		fooRet := foo.PerformSelector(Sel("methodForFoo"))
		if ToGoString(fooRet.Ptr()) != "foo" {
			t.Fatal("unexpected return")
		}

		bar := BarObjectClass.New()
		// bar go method
		if bar.MethodForBar() != "bar" {
			t.Fatal("unexpected return")
		}
		// bar objc method
		barRet := bar.PerformSelector(Sel("methodForBar"))
		if ToGoString(barRet.Ptr()) != "bar" {
			t.Fatal("unexpected return")
		}
		// as well as foo go method
		if bar.MethodForFoo() != "foo" {
			t.Fatal("unexpected return")
		}
		// and foo objc method
		barRet = bar.PerformSelector(Sel("methodForFoo"))
		if ToGoString(barRet.Ptr()) != "foo" {
			t.Fatal("unexpected return")
		}
	})

}

func TestNewClassNoMethod(t *testing.T) {
	assert.Panics(t, func() {
		NewClass[FooObject](
			Sel("noMethodLikeThisOnFooObject"),
		)
	})
}
