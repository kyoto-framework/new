package main

import "git.sr.ht/~kyoto-framework/kyoto"

// PIndexState is a state of PIndex.
type PIndexState struct {
	Example *kyoto.ComponentF[CExampleState]
}

// PIndex is a home page.
func PIndex(ctx *kyoto.Context) (state PIndexState) {
	// Setup rendering
	kyoto.Template(ctx, "page.index.go.html")
	// Init components
	state.Example = kyoto.Use(ctx, CExample(&CExampleArgs{
		Text: `
			This project includes
			tailwindcss as a fast prototyping tool,
			esbuild to build your custom js/css,
			zen for a comfortable work with Go
			and kyoto itself as a basis.
		`,
	}))
	// Return
	return
}
