package main

import (
	"net/http"

	"git.sr.ht/~kyoto-framework/kyoto"
	"git.sr.ht/~kyoto-framework/zen"
)

// setupStatic registers a static files handler.
func setupStatic() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./dist"))))
}

// setupKyoto provides advanced configuration for kyoto.
func setupKyoto() {
	kyoto.TemplateConf.FuncMap = kyoto.ComposeFuncMap(
		kyoto.FuncMap,
		zen.FuncMap,
	)
}

// setupPages registers project pages.
func setupPages() {
	// We are using kyoto.HandlePage as a simple way to register our pages.
	kyoto.HandlePage("/", PIndex)
	// Also, you can use kyoto.HandlerPage to create a raw http.HandlerFunc.
	// It might be useful for cases when you need to integrate
	// handler with another functions, like middlewares.
}

// setupActions registers actions for dynamic components.
func setupActions() {
	kyoto.HandleAction(CExample(nil))
}

// main is a project entry point.
func main() {
	setupStatic()
	setupKyoto()
	setupPages()
	setupActions()

	kyoto.Serve(":8000")
}
