package component

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/maxence-charriere/go-app/v10/pkg/ui"
	"osio88.github.io/src/internal/api"
)

type navigator struct {
	app.Compo

	streamInfo *api.StreamInfo
}

func newNavigator() *navigator {
	return &navigator{
		// TODO: make api function
		streamInfo: api.GetStream(50, 1, "viewer", "en"),
	}
}

func (c *navigator) Render() app.UI {
	return app.Div().
		Class("nav").
		Class("fill").
		Class("unselectable").
		// Class(n.Iclass).
		Body(
			ui.Stack().Class("app-title").Class("hspace-out").Middle().Content(
				newHeader(),
			),

			app.Nav().Class("nav-radios").Body(
				ui.Stack().Class("nav-radios-stack").Middle().Content(
					app.Div().Class("hspace-out").Body(
						app.Range(c.streamInfo.StreamList).Slice(func(i int) app.UI {
							return newLink(c.streamInfo.StreamList[i].ChannelID)
						}),
					),
				),
			),
		)
}
