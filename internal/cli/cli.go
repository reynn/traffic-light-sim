package cli

import (
	"flag"
	"time"
)

type (
	// Opts represents the command-line options for the traffic light simulator.
	Opts struct {
		YellowDuration time.Duration
		RedDuration    time.Duration
		GreenDuration  time.Duration
	}
)

// New returns a new instance of Opts with parsed values
func New() *Opts {
	opts := &Opts{}

	flag.DurationVar(&opts.YellowDuration, "yellow", 2*time.Second, "Duration of the yellow light")
	flag.DurationVar(&opts.RedDuration, "red", 3*time.Second, "Duration of the red light")
	flag.DurationVar(&opts.GreenDuration, "green", 3*time.Second, "Duration of the green light")

	flag.Parse()
	return opts
}
