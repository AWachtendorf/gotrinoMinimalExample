package components

import (
	"github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
	_ "github.com/golangee/gotrino-tailwind"
)

// AnotherComponent is a component that includes other components.
type AnotherComponent struct {
	gotrino.View
	FooComponent, BarComponent *SomeComponent
}

// NewAnotherComponent returns a new component and initializes the included components.
func NewAnotherComponent(string1, string2 string) *AnotherComponent {
	foo := NewComponent(string1)
	bar := NewComponent(string2)

	return &AnotherComponent{
		View:         gotrino.View{},
		FooComponent: foo,
		BarComponent: bar,
	}
}

// Render renders our node.
func (a *AnotherComponent) Render() gotrino.Node {

	return Div(
		Class("pt-10 bg-gray-400"),
		a.FooComponent,

		a.BarComponent,
		)
}
