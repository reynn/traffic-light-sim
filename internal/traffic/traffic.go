package traffic

import (
	"context"
	"fmt"
	"time"
)

type Light string

const (
	LightRed    Light = "red"
	LightYellow Light = "yellow"
	LightGreen  Light = "green"

	ASCIILightTop = `
       ##
      _[]_
     [____]
.----'    '----.`
	ASCIILightRedActive    = "|  \033[31;1m   .==.   \033[0m  |\n|  \033[31;1m  /    \\  \033[0m  |\n|  \033[31;1m  \\    /  \033[0m  |\n|  \033[31;1m   .==.   \033[0m  |"
	ASCIILightGreenActive  = "|  \033[32;1m   .==.   \033[0m  |\n|  \033[32;1m  /    \\  \033[0m  |\n|  \033[32;1m  \\    /  \033[0m  |\n|  \033[32;1m   .==.   \033[0m  |"
	ASCIILightYellowActive = "|  \033[33;1m   .==.   \033[0m  |\n|  \033[33;1m  /    \\  \033[0m  |\n|  \033[33;1m  \\    /  \033[0m  |\n|  \033[33;1m   .==.   \033[0m  |"
	ASCIILightInactive     = `|     .==.     |
|    /    \    |
|    \    /    |
|     .==.     |`
	ASCIILightBottom = `'--.________.--'`
)

type (
	LightController struct {
		RedLight    time.Duration
		YellowLight time.Duration
		GreenLight  time.Duration
	}
)

// NewLightController creates a new traffic light controller with the given durations
func NewLightController(red time.Duration, yellow time.Duration, green time.Duration) *LightController {
	return &LightController{
		RedLight:    red,
		YellowLight: yellow,
		GreenLight:  green,
	}
}

// Start Run the traffic light simulation
func (c *LightController) Start(ctx context.Context) error {
	ticker := time.NewTicker(c.RedLight)
	currLight := LightRed
	displayTrafficLight(currLight)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			var duration time.Duration
			currLight, duration = c.switchLight(currLight)
			ticker = time.NewTicker(duration)
			displayTrafficLight(currLight)
		}
	}
}

func (c *LightController) switchLight(currLight Light) (Light, time.Duration) {
	switch currLight {
	case LightRed:
		return LightYellow, c.YellowLight
	case LightYellow:
		return LightGreen, c.GreenLight
	case LightGreen:
		return LightRed, c.RedLight
	}
	return LightRed, c.RedLight
}

func displayTrafficLight(l Light) {
	fmt.Print("\033[H\033[2J")
	fmt.Println(ASCIILightTop)
	if l == LightRed {
		fmt.Println(ASCIILightRedActive) // red
	} else {
		fmt.Println(ASCIILightInactive) // red
	}
	if l == LightYellow {
		fmt.Println(ASCIILightYellowActive) // red
	} else {
		fmt.Println(ASCIILightInactive) // yellow
	}
	if l == LightGreen {
		fmt.Println(ASCIILightGreenActive) // red
	} else {
		fmt.Println(ASCIILightInactive) // yellow
	}
	fmt.Println(ASCIILightBottom)
	fmt.Println("")
	fmt.Println("Press ESC to quit")
}
