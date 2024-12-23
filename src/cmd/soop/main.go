package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"osio88.github.io/src/internal/component"
)

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {
	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", func() app.Composer { return component.NewSoopLive() })

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite("docs", &app.Handler{
		Name:        "Soooooop",
		Description: "This is Toy Soooooop",
		Resources:   app.GitHubPages("/"),
		Domain:      "osio88.github.io",
		Title:       "Soooooop",
		Image:       "/web/assets/images/icon.png",
		Icon: app.Icon{
			SVG:      "/web/assets/images/favicon.svg",
			Default:  "/web/assets/images/icon.png",
			Large:    "/web/assets/images/icon.png",
			Maskable: "/web/assets/images/icon.png",
		},
		Styles: []string{
			"https://fonts.googleapis.com/css2?family=Montserrat:wght@400;500&display=swap",
			"/web/assets/css/soop.css",
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	http.Handle("/", &app.Handler{
		Name:        "Soooooop",
		Description: "This is Toy Soooooop",
		// Resources:   app.GitHubPages("/"),
		Domain: "osio88.github.io",
		Title:  "Soooooop",
		Image:  "/web/assets/images/icon.png",
		Icon: app.Icon{
			SVG:      "/web/assets/images/favicon.svg",
			Default:  "/web/assets/images/icon.png",
			Large:    "/web/assets/images/icon.png",
			Maskable: "/web/assets/images/icon.png",
		},
		Styles: []string{
			"https://fonts.googleapis.com/css2?family=Montserrat:wght@400;500&display=swap",
			"/web/assets/css/soop.css",
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
