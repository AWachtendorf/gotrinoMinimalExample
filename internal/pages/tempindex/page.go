package tempindex

import (
	"github.com/golangee/dom/router"
	"github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
	"github.com/golangee/property"
)

const Path = "/"

type Index struct {
	gotrino.View

	previewModel property.String
}

func NewIndex() *Index {
	i := &Index{}
	i.previewModel.Attach(i.Invalidate)
	return i
}

// Render here implements the minimalistic start page with the search bar
func (i *Index) Render() gotrino.Node {
	return Div(
		Class("flex items-center justify-center"),
		Div(
			Text("Put some stuff in this body"),
			),
	)
}

func ApplyView(query router.Query) gotrino.Renderable {
	view := NewIndex()

	return view
}
