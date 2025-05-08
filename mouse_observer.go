package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

const (
	mouseTopic           = "mouse"
	mouseClick EventType = 1
)

func MouseObserver(bus *EventBus) {
	err := bus.NewTopic(mouseTopic)
	if err != nil {
		fmt.Printf("error creating topic: %s; topic might already exist\n", err)
	}

	go func() {
		prvMouseState := make(map[ebiten.MouseButton]bool)

		for {
			for _, btn := range []ebiten.MouseButton{
				ebiten.MouseButtonLeft,
				ebiten.MouseButtonRight,
				ebiten.MouseButtonMiddle,
			} {
				prev := prvMouseState[btn]
				curr := ebiten.IsMouseButtonPressed(btn)
				if !prev && curr {
					bus.Publish(mouseTopic,
						Event{
							ID:   mouseClick,
							Data: map[string]interface{}{"btn": btn, "pos": GetMousePos()},
						})
				}
				prvMouseState[btn] = curr
			}

			// Chatgpt told me that this was a ~120 FPS sleep system
			time.Sleep(8 * time.Millisecond)
		}
	}()
}
