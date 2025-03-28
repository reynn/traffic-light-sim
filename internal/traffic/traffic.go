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

func NewLightController(red time.Duration, yellow time.Duration, green time.Duration) *LightController {
	return &LightController{
		RedLight:    red,
		YellowLight: yellow,
		GreenLight:  green,
	}
}

func (c *LightController) Start(ctx context.Context) error {
	ticker := time.NewTicker(c.RedLight)
	currLight := LightRed
	displayTrafficLight(currLight)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			switch currLight {
			case LightRed:
				currLight = LightYellow
				ticker = time.NewTicker(c.YellowLight)
				displayTrafficLight(currLight)
			case LightYellow:
				currLight = LightGreen
				ticker = time.NewTicker(c.GreenLight)
				displayTrafficLight(currLight)
			case LightGreen:
				currLight = LightRed
				ticker = time.NewTicker(c.RedLight)
				displayTrafficLight(currLight)
			}
		}
	}
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
