package main

import (
	"github.com/joshjon/conduit-ci/sdk/conduit"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	w, err := conduit.NewWorker()
	if err != nil {
		return err
	}
	conduit.RegisterPipeline(w, Pipeline)
	conduit.RegisterArgResultJob(w, BarJob)
	return w.Run()
}
