package main

import (
	"context"
	"fmt"

	"github.com/joshjon/conduit-ci/sdk/conduit"
)

func Pipeline(ctx conduit.Context) error {
	_, err := conduit.StartArgResultJob(ctx, BarJob, BarConfig{Something: "foo"}).Wait(ctx)
	return err
}

type BarConfig struct {
	Something string
}

func BarJob(_ context.Context, cfg BarConfig) (string, error) {
	fmt.Println("bar job arg: ", cfg)
	return "answer", nil
}
