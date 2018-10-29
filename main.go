package main

import (
	"math/rand"
	"time"

	gobagcontext "github.com/danielkrainas/gobag/context"

	"github.com/tenderbytes/kindermud/pkg/cmd"
)

var appVersion string

const defaultVersion = "0.0.0-dev"

func main() {
	if appVersion == "" {
		appVersion = defaultVersion
	}

	rand.Seed(time.Now().Unix())

	ctx := gobagcontext.WithVersion(gobagcontext.Background(), appVersion)
	cmd.Execute(ctx)
}
