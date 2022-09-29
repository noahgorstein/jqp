package theme

import (
	"strings"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/styles"
	"github.com/charmbracelet/lipgloss"
)

const (
	BLUE  = lipgloss.Color("69")
	PINK  = lipgloss.Color("#F25D94")
	GREY  = lipgloss.Color("240")
	GREEN = lipgloss.Color("76")
	RED   = lipgloss.Color("9")
)

type Theme struct {
	Primary     lipgloss.Color
	Secondary   lipgloss.Color
	Inactive    lipgloss.Color
	Success     lipgloss.Color
	Error       lipgloss.Color
	ChromaStyle *chroma.Style
}

func getDefaultTheme() Theme {
	theme := Theme{
		Primary:     BLUE,
		Secondary:   PINK,
		Inactive:    GREY,
		Success:     GREEN,
		Error:       RED,
		ChromaStyle: styles.ParaisoLight,
	}
	if lipgloss.HasDarkBackground() {
		theme.ChromaStyle = styles.Vim
	}
	return theme
}

var (
	// from https://www.nordtheme.com/docs/colors-and-palettes
	nord7  = lipgloss.Color("#8FBCBB")
	nord9  = lipgloss.Color("#81A1C1")
	nord11 = lipgloss.Color("#BF616A")
	nord14 = lipgloss.Color("#A3BE8C")
)

var themeMap = map[string]Theme{
	"nord": {
		Primary:     nord7,
		Secondary:   nord9,
		Inactive:    GREY,
		Success:     nord14,
		Error:       nord11,
		ChromaStyle: styles.Nord,
	},
	"autumn": {
		Primary:     lipgloss.Color("#aa5500"),
		Secondary:   lipgloss.Color("#fcba03"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#009999"),
		Error:       lipgloss.Color("#ff0000"),
		ChromaStyle: styles.Autumn,
	},
	"monokai": {
		Primary:     lipgloss.Color("#a6e22e"),
		Secondary:   lipgloss.Color("#f92672"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#b4d273"),
		Error:       lipgloss.Color("#960050"),
		ChromaStyle: styles.Monokai,
	},
	"monokailight": {
		Primary:     lipgloss.Color("#00a8c8"),
		Secondary:   lipgloss.Color("#f92672"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#b4d273"),
		Error:       lipgloss.Color("#960050"),
		ChromaStyle: styles.MonokaiLight,
	},
	"dracula": {
		Primary:     lipgloss.Color("#8be9fd"),
		Secondary:   lipgloss.Color("#ffb86c"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#50fa7b"),
		Error:       lipgloss.Color("#f8f8f2"),
		ChromaStyle: styles.Dracula,
	},
}

func GetTheme(theme string) Theme {
	lowercasedTheme := strings.ToLower(strings.TrimSpace(theme))
	switch lowercasedTheme {
	case "default":
		return getDefaultTheme()
	case "nord":
		return themeMap["nord"]
	case "autumn":
		return themeMap["autumn"]
	case "monokai":
		return themeMap["monokai"]
	case "monokailight":
		return themeMap["monokailight"]
	case "dracula":
		return themeMap["dracula"]
	default:
		return getDefaultTheme()
	}
}
