package app

import (
	"github.com/golangee/dom/router"
	. "github.com/golangee/gotrino"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
	"rochus/frontend/internal/pages/tempindex"
)

type Application struct {
	router *router.Router
	log    log.Logger
}

func NewApplication() *Application {
	a := &Application{
		router: router.NewRouter(),
		log:    log.NewLogger(ecs.Log("application")),
	}

	// Add routes here. Just use the Path const you defined in your components.
	a.router.AddRoute(tempindex.Path, a.apply(tempindex.Render))

	return a
}

// apply is used, to render the content of
// the given function inside the body.
func (a *Application) apply(f func(query router.Query) Renderable) func(query router.Query) {
	return func(query router.Query) {
		RenderBody(f(query))
	}
}

func (a *Application) Run() {
	a.router.Start()
	select {}
}
