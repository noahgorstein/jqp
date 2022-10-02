package theme

import (
	"strings"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/styles"
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
	"abap": {
		Primary:     lipgloss.Color("#00f"),
		Secondary:   lipgloss.Color("#3af"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#5a2"),
		Error:       lipgloss.Color("#F00"),
		ChromaStyle: styles.Abap,
	},
	"algol": {
		Primary:     lipgloss.Color("#5a2"),
		Secondary:   lipgloss.Color("#666"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#5a2"),
		Error:       lipgloss.Color("#FF0000"),
		ChromaStyle: styles.Algol,
	},
	"arduino": {
		Primary:     lipgloss.Color("#1e90ff"),
		Secondary:   lipgloss.Color("#aa5500"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#5a2"),
		Error:       lipgloss.Color("#F00"),
		ChromaStyle: styles.Arduino,
	},
	"autumn": {
		Primary:     lipgloss.Color("#aa5500"),
		Secondary:   lipgloss.Color("#fcba03"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#009999"),
		Error:       lipgloss.Color("#ff0000"),
		ChromaStyle: styles.Autumn,
	},
	"average": {
		Primary:     lipgloss.Color("#ec0000"),
		Secondary:   lipgloss.Color("#008900"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#008900"),
		Error:       lipgloss.Color("#ec0000"),
		ChromaStyle: styles.Average,
	},
	"base16-snazzy": {
		Primary:     lipgloss.Color("#ff6ac1"),
		Secondary:   lipgloss.Color("#5af78e"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#5af78e"),
		Error:       lipgloss.Color("#ff5c57"),
		ChromaStyle: styles.Base16Snazzy,
	},
	"borland": {
		Primary:     lipgloss.Color("#00f"),
		Secondary:   lipgloss.Color("#000080"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#5a2"),
		Error:       lipgloss.Color("#a61717"),
		ChromaStyle: styles.Borland,
	},
	"colorful": {
		Primary:     lipgloss.Color("#00d"),
		Secondary:   lipgloss.Color("#070"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#070"),
		Error:       lipgloss.Color("#a61717"),
		ChromaStyle: styles.Colorful,
	},
	"doom-one": {
		Primary:     lipgloss.Color("#b756ff"),
		Secondary:   lipgloss.Color("#63c381"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#63c381"),
		Error:       lipgloss.Color("#e06c75"),
		ChromaStyle: styles.DoomOne,
	},
	"doom-one2": {
		Primary:     lipgloss.Color("#76a9f9"),
		Secondary:   lipgloss.Color("#63c381"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#63c381"),
		Error:       lipgloss.Color("#e06c75"),
		ChromaStyle: styles.DoomOne2,
	},
	"dracula": {
		Primary:     lipgloss.Color("#8be9fd"),
		Secondary:   lipgloss.Color("#ffb86c"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#50fa7b"),
		Error:       lipgloss.Color("#f8f8f2"),
		ChromaStyle: styles.Dracula,
	},
	"emacs": {
		Primary:     lipgloss.Color("#008000"),
		Secondary:   lipgloss.Color("#a2f"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#008000"),
		Error:       lipgloss.Color("#b44"),
		ChromaStyle: styles.Emacs,
	},
	"friendly": {
		Primary:     lipgloss.Color("#40a070"),
		Secondary:   lipgloss.Color("#062873"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#40a070"),
		Error:       lipgloss.Color("#FF0000"),
		ChromaStyle: styles.Friendly,
	},
	"fruity": {
		Primary:     lipgloss.Color("#fb660a"),
		Secondary:   lipgloss.Color("#0086f7"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#40a070"),
		Error:       lipgloss.Color("#FF0000"),
		ChromaStyle: styles.Fruity,
	},
	"github": {
		Primary:     lipgloss.Color("#d14"),
		Secondary:   lipgloss.Color("#099"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#099"),
		Error:       lipgloss.Color("#d14"),
		ChromaStyle: styles.GitHub,
	},
	"github-dark": {
		Primary:     lipgloss.Color("#d2a8ff"),
		Secondary:   lipgloss.Color("#f0883e"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#56d364"),
		Error:       lipgloss.Color("#ffa198"),
		ChromaStyle: styles.GitHubDark,
	},
	"gruvbox": {
		Primary:     lipgloss.Color("#b8bb26"),
		Secondary:   lipgloss.Color("#d3869b"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#b8bb26"),
		Error:       lipgloss.Color("#fb4934"),
		ChromaStyle: styles.Gruvbox,
	},
	"gruvbox-light": {
		Primary:     lipgloss.Color("#fb4934"),
		Secondary:   lipgloss.Color("#b8bb26"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#b8bb26"),
		Error:       lipgloss.Color("#9D0006"),
		ChromaStyle: styles.GruvboxLight,
	},
	"hrdark": {
		Primary:     lipgloss.Color("#58a1dd"),
		Secondary:   lipgloss.Color("#ff636f"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#a6be9d"),
		Error:       lipgloss.Color("#FF0000"),
		ChromaStyle: styles.HrDark,
	},
	"igor": {
		Primary:     lipgloss.Color("#009c00"),
		Secondary:   lipgloss.Color("#00f"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#009c00"),
		Error:       lipgloss.Color("#FF0000"),
		ChromaStyle: styles.Igor,
	},
	"lovelace": {
		Primary:     lipgloss.Color("#b83838"),
		Secondary:   lipgloss.Color("#2838b0"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#009c00"),
		Error:       lipgloss.Color("#b83838"),
		ChromaStyle: styles.Igor,
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
	"nord": {
		Primary:     nord7,
		Secondary:   nord9,
		Inactive:    GREY,
		Success:     nord14,
		Error:       nord11,
		ChromaStyle: styles.Nord,
	},
	"paradaiso-dark": {
		Primary:     lipgloss.Color("#48b685"),
		Secondary:   lipgloss.Color("#5bc4bf"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#48b685"),
		Error:       lipgloss.Color("#ef6155"),
		ChromaStyle: styles.ParaisoDark,
	},
	"paradaiso-light": {
		Primary:     lipgloss.Color("#48b685"),
		Secondary:   lipgloss.Color("#815ba4"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#48b685"),
		Error:       lipgloss.Color("#ef6155"),
		ChromaStyle: styles.ParaisoLight,
	},
	"pygments": {
		Primary:     lipgloss.Color("#008000"),
		Secondary:   lipgloss.Color("#ba2121"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#008000"),
		Error:       lipgloss.Color("#ba2121"),
		ChromaStyle: styles.Pygments,
	},
}

func GetTheme(theme string) Theme {
	lowercasedTheme := strings.ToLower(strings.TrimSpace(theme))
	if value, ok := themeMap[lowercasedTheme]; ok {
		return value
	} else {
		return getDefaultTheme()
	}
}
