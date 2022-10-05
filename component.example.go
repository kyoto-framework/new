package main

import (
	"git.sr.ht/~kyoto-framework/kyoto"
	"git.sr.ht/~kyoto-framework/zen"
)

// CExampleArgs is an arguments holder of CExample.
type CExampleArgs struct {
	Text string
}

// CExampleState is a state of CExample.
type CExampleState struct {
	Args *CExampleArgs

	UUID string `json:"-"` // Will be reloaded each time, no sense to keep in state
}

// CExample is an example component
// to show a preferred way of components definition in kyoto.
// Supposed to be deleted on project start.
func CExample(args *CExampleArgs) kyoto.Component[CExampleState] {
	return func(ctx *kyoto.Context) (state CExampleState) {
		// Write args
		state.Args = args
		// Preload action state
		kyoto.ActionPreload(ctx, &state)
		// Example data fetching
		data := map[string]string{}
		zen.Request("GET", "https://httpbin.org/uuid").
			Do().Success().Decode(&data).Must()
		// Set state from data
		state.UUID = data["uuid"]
		// Return
		return
	}
}
