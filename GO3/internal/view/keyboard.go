package view

import (
	"s21_rogue/internal/domain/value_objects"

	"github.com/gdamore/tcell/v3"
)

type Keyboard interface {
	GetActionChan() chan value_objects.Action
}

type tcellKeyboard struct {
	screen *tcell.Screen
}

func NewKeyboard(screen *tcell.Screen) Keyboard {
	return &tcellKeyboard{screen: screen}
}

func (k *tcellKeyboard) GetActionChan() chan value_objects.Action {
	c := make(chan value_objects.Action)

	go func() {
		defer close(c)
		for {
			ev := <-(*k.screen).EventQ()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				result := value_objects.Nothing
				switch ev.Key() {
				case tcell.KeyUp:
					result = value_objects.UpAction
				case tcell.KeyRight:
					result = value_objects.RightAction
				case tcell.KeyDown:
					result = value_objects.DownAction
				case tcell.KeyLeft:
					result = value_objects.LeftAction
				case tcell.KeyEsc:
					result = value_objects.TerminateAction
				case tcell.KeyEnter:
					result = value_objects.StartAction
				}

				if result == value_objects.Nothing {
					switch ev.Str() {
					case "w":
						result = value_objects.UpAction
					case "d":
						result = value_objects.RightAction
					case "s":
						result = value_objects.DownAction
					case "a":
						result = value_objects.LeftAction
					case "j":
						result = value_objects.FoodAction
					case "e":
						result = value_objects.ScrollAction
					case "k":
						result = value_objects.ElixirAction
					case "h":
						result = value_objects.WeaponAction
					default:
						result = value_objects.Nothing
					}
				}
				c <- result

			}
		}
	}()
	return c
}
