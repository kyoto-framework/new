package main

import (
	"net/http"

	"git.sr.ht/~kyoto-framework/kyoto"
)

func setupStatic() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./dist"))))
}

func setupPages() {
	kyoto.HandlePage("/", PIndex)
}

func setupActions() {
	kyoto.HandleAction(CExample(nil))
}

func main() {
	setupStatic()
	setupPages()
	setupActions()

	kyoto.Serve(":8000")
}
