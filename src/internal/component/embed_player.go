package component

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

const defaultChannelId = "sherlock"
const embedUrl = "https://www.sooplive.com/player/embed/%s"

type embedPlayer struct {
	app.Compo

	channelId string
}

func newEmbedPlayer() *embedPlayer {
	return &embedPlayer{}
}

func (c *embedPlayer) OnMount(ctx app.Context) {
	channelId := ""
	err := ctx.LocalStorage().Get("channelId", &channelId)

	if err != nil || channelId == "" {
		channelId = defaultChannelId
		ctx.LocalStorage().Set("channelId", channelId)
	}

	c.channelId = channelId
}

func (c *embedPlayer) Render() app.UI {
	return app.Div().
		Class("youtube").
		Body(
			app.Div().Class("youtube-video").Body(
				app.Div().ID("youtube-player").
					Class("unselectable").
					Body(
						app.IFrame().Src(fmt.Sprintf(embedUrl, c.channelId)),
					),
			),
		)
}
