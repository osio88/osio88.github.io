package component

import (
	"log/slog"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type SoopLive struct {
	app.Compo
}

func NewSoopLive() *SoopLive {
	return &SoopLive{}
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (c *SoopLive) Render() app.UI {
	slog.Info("Start Soooooooooop Live")
	return app.Main().Body(newNavigator(), newEmbedPlayer())
}
