package components

import (
	"github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
	_ "github.com/golangee/gotrino-tailwind" //sometimes the tailwind styles don't work, I just use the passive import.
	)

// SomeComponent is an example struct. Always include a View. The rest may be all kinds of variables.
type SomeComponent struct{
	gotrino.View
	content string
}

// NewComponent is a basic constructor that you may pass a string to.
func NewComponent(someText string)*SomeComponent{
	return &SomeComponent{
		View:    gotrino.View{},
		content: someText,
	}
}

// Render is always the same. It renders our nodes. You have a chance to style your components here.
func (s *SomeComponent)Render()gotrino.Node{

	return Div(
		Class("p-10 m-10 bg-gray-300"),Text(s.content),
		)

}