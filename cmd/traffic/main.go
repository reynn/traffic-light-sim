package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/eiannone/keyboard"

	"github.com/reynn/traffic-light-sim/internal/cli"
	"github.com/reynn/traffic-light-sim/internal/traffic"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	cliOpts := cli.New()

	controller := traffic.NewLightController(
		cliOpts.RedDuration,
		cliOpts.YellowDuration,
		cliOpts.GreenDuration,
	)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		if e := controller.Start(ctx); e != nil && !errors.Is(e, context.Canceled) {
			fmt.Printf("Error with light controller: %v\n", e)
		}
	}()
	// TODO: Add keyboard control for changing traffic light durations during operation
	go func() {
		defer wg.Done()
		keysEvents, err := keyboard.GetKeys(10)
		if err != nil {
			panic(err)
		}
		defer func() {
			if e := keyboard.Close(); e != nil {
				fmt.Printf("Error closing keyboard: %v\n", e)
			}
		}()

		for {
			event := <-keysEvents
			if event.Key == keyboard.KeyEsc {
				cancel()
				break
			}
		}
	}()

	wg.Wait()
	fmt.Println("Exiting Traffic Light Simulator")
}
