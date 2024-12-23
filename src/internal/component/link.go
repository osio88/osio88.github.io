package component

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/maxence-charriere/go-app/v10/pkg/ui"
)

type link struct {
	app.Compo

	channelId string
	href      string

	IonClick func(ctx app.Context)
}

const embedLink = "https://www.sooplive.com/player/embed/%s"

func newLink(channelId string) *link {
	return &link{
		channelId: channelId,
		href:      fmt.Sprintf(embedLink, channelId),
	}
}

func (c *link) onClick(ctx app.Context, e app.Event) {
	e.PreventDefault()
	ctx.LocalStorage().Set("channelId", &c.channelId)
	ctx.Reload()
}

func (c *link) Render() app.UI {
	return app.A().
		ID(c.channelId).
		Class("link").
		Class("heading").
		Class("fit").
		Title(c.channelId).
		Href(c.href).
		OnClick(c.onClick).
		Body(
			ui.Stack().Middle().Content(
				app.Div().Class("link-icon"),
				app.Div().Text(c.channelId),
			),
		)
}
