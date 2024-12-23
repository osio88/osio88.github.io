package component

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type header struct {
	app.Compo
}

func newHeader() *header {
	return &header{}
}

func (c *header) Render() app.UI {
	return app.Header().Body(
		app.A().
			Class("hApp").
			Class("focus").
			Class("glow").
			Href("/").
			Text("SoopLive"),
	)
}
