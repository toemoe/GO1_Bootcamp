package view

import (
	"github.com/gdamore/tcell/v3"
	"github.com/gdamore/tcell/v3/color"
)

type Theme struct {
	Background tcell.Color

	Map        MapTheme
	Characters CharactersTheme
	Items      ItemsTheme
	HUD        HudTheme
	MainMenu   MainMenuTheme
}

type MapTheme struct {
	Wall       tcell.Color
	Current    tcell.Color
	Visited    tcell.Color
	NotVisited tcell.Color
	Corridor   tcell.Color
	Portal     tcell.Color
}

type ItemsTheme struct {
	Treasure tcell.Color
	Food     tcell.Color
	Elixir   tcell.Color
	Scroll   tcell.Color
	Weapon   tcell.Color
}

type CharactersTheme struct {
	Player  tcell.Color
	Zombie  tcell.Color
	Ghost   tcell.Color
	Vampire tcell.Color
	Ogre    tcell.Color
	Snake   tcell.Color
}

type HudTheme struct {
	Text    tcell.Color
	Value   tcell.Color
	Warning tcell.Color
}

type MainMenuTheme struct {
	H1Text tcell.Color
}

type ViewerConfig struct {
	Theme Theme
}

func BasicViewerConfig() ViewerConfig {
	return ViewerConfig{
		Theme: Theme{
			Background: color.Black,

			Map: MapTheme{
				Wall:       color.Black,
				Current:    color.DarkRed,
				Visited:    color.DarkGray,
				NotVisited: color.Black,
				Portal:     color.DarkTurquoise,
			},

			Items: ItemsTheme{
				Food:     color.Orange,
				Treasure: color.LightGoldenrodYellow,
				Elixir:   color.Red,
				Scroll:   color.Gold,
				Weapon:   color.LightSlateGray,
			},

			Characters: CharactersTheme{
				Player:  color.DarkGray,
				Zombie:  color.DarkGreen,
				Ghost:   color.DarkBlue,
				Vampire: color.DarkMagenta,
				Ogre:    color.DarkKhaki,
				Snake:   color.DarkViolet,
			},

			HUD: HudTheme{
				Text:    color.Gray,
				Value:   color.Gold,
				Warning: color.Red,
			},

			MainMenu: MainMenuTheme{
				H1Text: color.DarkRed,
			},
		},
	}
}

type ViewerStyles struct {
	MapStyles  MapStyles
	Characters CharactersStyles
	Items      ItemsStyles
	Ui         Ui
}

type MapStyles struct {
	Wall       tcell.Style
	Current    tcell.Style
	Visited    tcell.Style
	NotVisited tcell.Style
	Portal     tcell.Style
}

type ItemsStyles struct {
	Food     tcell.Style
	Treasure tcell.Style
	Elixir   tcell.Style
	Scroll   tcell.Style
	Weapon   tcell.Style
}

type CharactersStyles struct {
	Player  tcell.Style
	Zombie  tcell.Style
	Ghost   tcell.Style
	Vampire tcell.Style
	Ogre    tcell.Style
	Snake   tcell.Style
}

type Ui struct {
	Text     tcell.Style
	Warning  tcell.Style
	Vignette tcell.Style
	Value    tcell.Style
	H1       tcell.Style
}

func BasicViewerStyles() ViewerStyles {
	config := BasicViewerConfig()
	return ViewerStyles{
		MapStyles: MapStyles{
			Wall: tcell.StyleDefault.
				Foreground(config.Theme.Map.Wall).
				Background(config.Theme.Background),

			Current: tcell.StyleDefault.
				Foreground(config.Theme.Map.Current).
				Background(config.Theme.Background),

			NotVisited: tcell.StyleDefault.
				Foreground(config.Theme.Map.NotVisited).
				Background(config.Theme.Background),

			Visited: tcell.StyleDefault.
				Foreground(config.Theme.Map.Visited).
				Background(config.Theme.Background),

			Portal: tcell.StyleDefault.
				Foreground(config.Theme.Map.Portal).
				Background(config.Theme.Background),
		},

		Ui: Ui{
			Text: tcell.StyleDefault.
				Foreground(config.Theme.HUD.Text).
				Background(config.Theme.Background),

			Warning: tcell.StyleDefault.
				Foreground(config.Theme.HUD.Warning).
				Background(config.Theme.Background),

			Value: tcell.StyleDefault.
				Foreground(config.Theme.HUD.Value).
				Background(config.Theme.Background),

			H1: tcell.StyleDefault.Bold(true).
				Foreground(config.Theme.MainMenu.H1Text).
				Background(config.Theme.Background),
			Vignette: tcell.StyleDefault.
				Foreground(config.Theme.Background).
				Background(config.Theme.Background),
		},

		Items: ItemsStyles{
			Food: tcell.StyleDefault.
				Foreground(config.Theme.Items.Food).
				Background(config.Theme.Background),

			Treasure: tcell.StyleDefault.
				Foreground(config.Theme.Items.Treasure).
				Background(config.Theme.Background),

			Weapon: tcell.StyleDefault.
				Foreground(config.Theme.Items.Weapon).
				Background(config.Theme.Background),

			Elixir: tcell.StyleDefault.
				Foreground(config.Theme.Items.Elixir).
				Background(config.Theme.Background),

			Scroll: tcell.StyleDefault.
				Foreground(config.Theme.Items.Scroll).
				Background(config.Theme.Background),
		},

		Characters: CharactersStyles{
			Player: tcell.StyleDefault.
				Foreground(config.Theme.Characters.Player).
				Background(config.Theme.Background),

			Zombie: tcell.StyleDefault.
				Foreground(config.Theme.Characters.Zombie).
				Background(config.Theme.Background),

			Vampire: tcell.StyleDefault.
				Foreground(config.Theme.Characters.Vampire).
				Background(config.Theme.Background),

			Snake: tcell.StyleDefault.
				Foreground(config.Theme.Characters.Snake).
				Background(config.Theme.Background),

			Ogre: tcell.StyleDefault.
				Foreground(config.Theme.Characters.Ogre).
				Background(config.Theme.Background),

			Ghost: tcell.StyleDefault.
				Foreground(config.Theme.Characters.Ghost).
				Background(config.Theme.Background),
		},
	}
}
