package app

import (
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/content"
	"github.com/leandro-grg/go-api/app/handlers"
)

// CmdServe - Serves the app
func CmdServe() {
	router := routing.New()

	api := router.Group("/api")
	api.Use(content.TypeNegotiator(content.JSON))

	defaultHandler := handlers.DefaultHandler{}

	api.Get("/default", defaultHandler.Index)

	http.Handle("/", router)
	panic(http.ListenAndServe("localhost:8080", nil))
}
