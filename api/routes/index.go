package routes

import core "learn.zone01dakar.sn/forum-rest-api/internals/core"

type Router interface {
	Route(app *core.App)
}

func Handle(a []Router, app *core.App) {
	for _, s := range a {
		s.Route(app)
	}
}
