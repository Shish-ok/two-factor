package main

import (
	"go.uber.org/fx"
	"two-factor-auth/internal/core"
)

// @title two-factor-auth doc
// main godoc
// @in header
func main() {
	fx.New(core.Core()).Run()
}
