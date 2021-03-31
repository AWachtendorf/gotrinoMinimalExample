package tempindex

import (
	"github.com/golangee/dom/router"
	"github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
	_ "github.com/golangee/gotrino-tailwind"
	"rochus/frontend/internal/components"
)

const Path = "/"

// Render renders our content via query.
func Render(query router.Query) gotrino.Renderable {

	// initialize some components.
	Foo := components.NewComponent("I am a text.")
	Bar := components.NewComponent("I'm also a text. Same Component but different content")
	FooBar := components.NewAnotherComponent("This is a text in a component", "This also")

	// returns our content.
	return Div(
		// just put in some style.
		Class("bg-gray-200"),
		// this div now is the place for all our useless components
		Div(
			Foo, // render the Foo component.
			Bar, // render the Bar component.
			FooBar, // Render the Foobar component, that also includes a Foo and a Bar Component.
			Bar, // Render BAr component again, just to show reuseablility.
		),
	)
}
