package theme

import (
	"strings"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/charmbracelet/lipgloss"
)

type CustomTheme struct {
	Primary   string
	Secondary string
	Inactive  string
	Success   string
	Error     string
}

var CustomThemeKeys = CustomTheme{
	Primary:   "primary",
	Secondary: "secondary",
	Success:   "success",
	Inactive:  "inactive",
	Error:     "error",
}

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
		ChromaStyle: styles.Get("paradaiso-dark"),
	}
	if lipgloss.HasDarkBackground() {
		theme.ChromaStyle = styles.Get("vim")
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
		ChromaStyle: styles.Get("abap"),
	},
	"algol": {
		Primary:     lipgloss.Color("#5a2"),
		Secondary:   lipgloss.Color("#666"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#5a2"),
		Error:       lipgloss.Color("#FF0000"),
		ChromaStyle: styles.Get("algol"),
	},
	"arduino": {
		Primary:     lipgloss.Color("#1e90ff"),
		Secondary:   lipgloss.Color("#aa5500"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#5a2"),
		Error:       lipgloss.Color("#F00"),
		ChromaStyle: styles.Get("arduino"),
	},
	"autumn": {
		Primary:     lipgloss.Color("#aa5500"),
		Secondary:   lipgloss.Color("#fcba03"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#009999"),
		Error:       lipgloss.Color("#ff0000"),
		ChromaStyle: styles.Get("autumn"),
	},
	"average": {
		Primary:     lipgloss.Color("#ec0000"),
		Secondary:   lipgloss.Color("#008900"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#008900"),
		Error:       lipgloss.Color("#ec0000"),
		ChromaStyle: styles.Get("average"),
	},
	"base16-snazzy": {
		Primary:     lipgloss.Color("#ff6ac1"),
		Secondary:   lipgloss.Color("#5af78e"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#5af78e"),
		Error:       lipgloss.Color("#ff5c57"),
		ChromaStyle: styles.Get("base16-snazzy"),
	},
	"borland": {
		Primary:     lipgloss.Color("#00f"),
		Secondary:   lipgloss.Color("#000080"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#5a2"),
		Error:       lipgloss.Color("#a61717"),
		ChromaStyle: styles.Get("borland"),
	},
	"catppuccin-latte": {
		Primary:     lipgloss.Color("#179299"),
		Secondary:   lipgloss.Color("#1e66f5"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#40a02b"),
		Error:       lipgloss.Color("#d20f39"),
		ChromaStyle: styles.Get("catppuccin-latte"),
	},
	"catppuccin-frappe": {
		Primary:     lipgloss.Color("#81c8be"),
		Secondary:   lipgloss.Color("#8caaee"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#a6d189"),
		Error:       lipgloss.Color("#e78284"),
		ChromaStyle: styles.Get("catppuccin-frappe"),
	},
	"catppuccin-macchiato": {
		Primary:     lipgloss.Color("#8bd5ca"),
		Secondary:   lipgloss.Color("#8aadf4"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#a6da95"),
		Error:       lipgloss.Color("#ed8796"),
		ChromaStyle: styles.Get("catppuccin-macchiato"),
	},
	"catppuccin-mocha": {
		Primary:     lipgloss.Color("#94e2d5"),
		Secondary:   lipgloss.Color("#89b4fa"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#a6e3a1"),
		Error:       lipgloss.Color("#f38ba8"),
		ChromaStyle: styles.Get("catppuccin-mocha"),
	},
	"colorful": {
		Primary:     lipgloss.Color("#00d"),
		Secondary:   lipgloss.Color("#070"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#070"),
		Error:       lipgloss.Color("#a61717"),
		ChromaStyle: styles.Get("colorful"),
	},
	"doom-one": {
		Primary:     lipgloss.Color("#b756ff"),
		Secondary:   lipgloss.Color("#63c381"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#63c381"),
		Error:       lipgloss.Color("#e06c75"),
		ChromaStyle: styles.Get("doom-one"),
	},
	"doom-one2": {
		Primary:     lipgloss.Color("#76a9f9"),
		Secondary:   lipgloss.Color("#63c381"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#63c381"),
		Error:       lipgloss.Color("#e06c75"),
		ChromaStyle: styles.Get("doom-one2"),
	},
	"dracula": {
		Primary:     lipgloss.Color("#8be9fd"),
		Secondary:   lipgloss.Color("#ffb86c"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#50fa7b"),
		Error:       lipgloss.Color("#f8f8f2"),
		ChromaStyle: styles.Get("dracula"),
	},
	"emacs": {
		Primary:     lipgloss.Color("#008000"),
		Secondary:   lipgloss.Color("#a2f"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#008000"),
		Error:       lipgloss.Color("#b44"),
		ChromaStyle: styles.Get("emacs"),
	},
	"friendly": {
		Primary:     lipgloss.Color("#40a070"),
		Secondary:   lipgloss.Color("#062873"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#40a070"),
		Error:       lipgloss.Color("#FF0000"),
		ChromaStyle: styles.Get("friendly"),
	},
	"fruity": {
		Primary:     lipgloss.Color("#fb660a"),
		Secondary:   lipgloss.Color("#0086f7"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#40a070"),
		Error:       lipgloss.Color("#FF0000"),
		ChromaStyle: styles.Get("fruity"),
	},
	"github": {
		Primary:     lipgloss.Color("#d14"),
		Secondary:   lipgloss.Color("#099"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#099"),
		Error:       lipgloss.Color("#d14"),
		ChromaStyle: styles.Get("github"),
	},
	"github-dark": {
		Primary:     lipgloss.Color("#d2a8ff"),
		Secondary:   lipgloss.Color("#f0883e"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#56d364"),
		Error:       lipgloss.Color("#ffa198"),
		ChromaStyle: styles.Get("github-dark"),
	},
	"gruvbox": {
		Primary:     lipgloss.Color("#b8bb26"),
		Secondary:   lipgloss.Color("#d3869b"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#b8bb26"),
		Error:       lipgloss.Color("#fb4934"),
		ChromaStyle: styles.Get("gruvbox"),
	},
	"gruvbox-light": {
		Primary:     lipgloss.Color("#fb4934"),
		Secondary:   lipgloss.Color("#b8bb26"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#b8bb26"),
		Error:       lipgloss.Color("#9D0006"),
		ChromaStyle: styles.Get("gruvbox-light"),
	},
	"hrdark": {
		Primary:     lipgloss.Color("#58a1dd"),
		Secondary:   lipgloss.Color("#ff636f"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#a6be9d"),
		Error:       lipgloss.Color("#FF0000"),
		ChromaStyle: styles.Get("hrdark"),
	},
	"igor": {
		Primary:     lipgloss.Color("#009c00"),
		Secondary:   lipgloss.Color("#00f"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#009c00"),
		Error:       lipgloss.Color("#FF0000"),
		ChromaStyle: styles.Get("igor"),
	},
	"lovelace": {
		Primary:     lipgloss.Color("#b83838"),
		Secondary:   lipgloss.Color("#2838b0"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#009c00"),
		Error:       lipgloss.Color("#b83838"),
		ChromaStyle: styles.Get("lovelace"),
	},
	"manni": {
		Primary:     lipgloss.Color("#c30"),
		Secondary:   lipgloss.Color("#309"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#009c00"),
		Error:       lipgloss.Color("#c30"),
		ChromaStyle: styles.Get("manni"),
	},
	"monokai": {
		Primary:     lipgloss.Color("#a6e22e"),
		Secondary:   lipgloss.Color("#f92672"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#b4d273"),
		Error:       lipgloss.Color("#960050"),
		ChromaStyle: styles.Get("monokai"),
	},
	"monokai-light": {
		Primary:     lipgloss.Color("#00a8c8"),
		Secondary:   lipgloss.Color("#f92672"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#b4d273"),
		Error:       lipgloss.Color("#960050"),
		ChromaStyle: styles.Get("monokai-light"),
	},
	"murphy": {
		Primary:     lipgloss.Color("#070"),
		Secondary:   lipgloss.Color("#66f"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#070"),
		Error:       lipgloss.Color("#F00"),
		ChromaStyle: styles.Get("murphy"),
	},
	"native": {
		Primary:     lipgloss.Color("#6ab825"),
		Secondary:   lipgloss.Color("#ed9d13"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#6ab825"),
		Error:       lipgloss.Color("#a61717"),
		ChromaStyle: styles.Get("native"),
	},
	"nord": {
		Primary:     nord7,
		Secondary:   nord9,
		Inactive:    GREY,
		Success:     nord14,
		Error:       nord11,
		ChromaStyle: styles.Get("nord"),
	},
	"onesenterprise": {
		Primary:     lipgloss.Color("#00f"),
		Secondary:   lipgloss.Color("#f00"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#6ab825"),
		Error:       lipgloss.Color("#f00"),
		ChromaStyle: styles.Get("onesenterprise"),
	},
	"pastie": {
		Primary:     lipgloss.Color("#b06"),
		Secondary:   lipgloss.Color("#00d"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#080"),
		Error:       lipgloss.Color("#d20"),
		ChromaStyle: styles.Get("pastie"),
	},
	"perldoc": {
		Primary:     lipgloss.Color("#8b008b"),
		Secondary:   lipgloss.Color("#b452cd"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#080"),
		Error:       lipgloss.Color("#cd5555"),
		ChromaStyle: styles.Get("perldoc"),
	},
	"paradaiso-dark": {
		Primary:     lipgloss.Color("#48b685"),
		Secondary:   lipgloss.Color("#5bc4bf"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#48b685"),
		Error:       lipgloss.Color("#ef6155"),
		ChromaStyle: styles.Get("paradaiso-dark"),
	},
	"paradaiso-light": {
		Primary:     lipgloss.Color("#48b685"),
		Secondary:   lipgloss.Color("#815ba4"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#48b685"),
		Error:       lipgloss.Color("#ef6155"),
		ChromaStyle: styles.Get("paradaiso-light"),
	},
	"pygments": {
		Primary:     lipgloss.Color("#008000"),
		Secondary:   lipgloss.Color("#ba2121"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#008000"),
		Error:       lipgloss.Color("#ba2121"),
		ChromaStyle: styles.Get("pygments"),
	},
	"rainbow_dash": {
		Primary:     lipgloss.Color("#0c6"),
		Secondary:   lipgloss.Color("#5918bb"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#0c6"),
		Error:       lipgloss.Color("#ba2121"),
		ChromaStyle: styles.Get("rainbow_dash"),
	},
	"rrt": {
		Primary:     lipgloss.Color("#f60"),
		Secondary:   lipgloss.Color("#87ceeb"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#0c6"),
		Error:       lipgloss.Color("#f00"),
		ChromaStyle: styles.Get("rrt"),
	},
	"solarized-dark": {
		Primary:     lipgloss.Color("#268bd2"),
		Secondary:   lipgloss.Color("#2aa198"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#0c6"),
		Error:       lipgloss.Color("#cb4b16"),
		ChromaStyle: styles.Get("solarized-dark"),
	},
	"solarized-dark256": {
		Primary:     lipgloss.Color("#0087ff"),
		Secondary:   lipgloss.Color("#00afaf"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#0c6"),
		Error:       lipgloss.Color("#d75f00"),
		ChromaStyle: styles.Get("solarized-dark256"),
	},
	"solarized-light": {
		Primary:     lipgloss.Color("#268bd2"),
		Secondary:   lipgloss.Color("#2aa198"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#859900"),
		Error:       lipgloss.Color("#d75f00"),
		ChromaStyle: styles.Get("solarized-light"),
	},
	"swapoff": {
		Primary:     lipgloss.Color("#0ff"),
		Secondary:   lipgloss.Color("#ff0"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#e5e5e5"),
		Error:       lipgloss.Color("#e5e5e5"),
		ChromaStyle: styles.Get("swapoff"),
	},
	"tango": {
		Primary:     lipgloss.Color("#204a87"),
		Secondary:   lipgloss.Color("#0000cf"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#4e9a06"),
		Error:       lipgloss.Color("#a40000"),
		ChromaStyle: styles.Get("tango"),
	},
	"trac": {
		Primary:     lipgloss.Color("#099"),
		Secondary:   lipgloss.Color("#000080"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#099"),
		Error:       lipgloss.Color("#a61717"),
		ChromaStyle: styles.Get("trac"),
	},
	"vim": {
		Primary:     lipgloss.Color("#cd00cd"),
		Secondary:   lipgloss.Color("#cdcd00"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#66FF00"),
		Error:       lipgloss.Color("#cd0000"),
		ChromaStyle: styles.Get("vim"),
	},
	"visual_studio": {
		Primary:     lipgloss.Color("#a31515"),
		Secondary:   lipgloss.Color("#00f"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#023020"),
		Error:       lipgloss.Color("#a31515"),
		ChromaStyle: styles.Get("vs"),
	},
	"vulcan": {
		Primary:     lipgloss.Color("#bc74c4"),
		Secondary:   lipgloss.Color("#56b6c2"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#82cc6a"),
		Error:       lipgloss.Color("#cf5967"),
		ChromaStyle: styles.Get("vulcan"),
	},
	"witchhazel": {
		Primary:     lipgloss.Color("#ffb8d1"),
		Secondary:   lipgloss.Color("#56b6c2"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#c2ffdf"),
		Error:       lipgloss.Color("#ffb8d1"),
		ChromaStyle: styles.Get("witchhazel"),
	},
	"xcode": {
		Primary:     lipgloss.Color("#c41a16"),
		Secondary:   lipgloss.Color("#1c01ce"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#023020"),
		Error:       lipgloss.Color("#c41a16"),
		ChromaStyle: styles.Get("xcode"),
	},
	"xcode-dark": {
		Primary:     lipgloss.Color("#fc6a5d"),
		Secondary:   lipgloss.Color("#d0bf69"),
		Inactive:    GREY,
		Success:     lipgloss.Color("#90EE90"),
		Error:       lipgloss.Color("#fc6a5d"),
		ChromaStyle: styles.Get("xcode-dark"),
	},
}

// returns a theme by name, and true if default theme was returned
func GetTheme(themeName string, styleOverrides map[string]string) (Theme, bool) {
	lowercasedTheme := strings.ToLower(strings.TrimSpace(themeName))

	var isDefault bool
	var theme Theme
	if value, ok := themeMap[lowercasedTheme]; ok {
		theme = value
		isDefault = false
	} else {
		theme = getDefaultTheme()
		isDefault = true
	}

	theme.SetOverrides(styleOverrides)

	return theme, isDefault && len(styleOverrides) == 0
}

func (t *Theme) SetOverrides(overrides map[string]string) {
	t.Primary = customColorOrDefault(overrides[CustomThemeKeys.Primary], t.Primary)
	t.Secondary = customColorOrDefault(overrides[CustomThemeKeys.Secondary], t.Secondary)
	t.Inactive = customColorOrDefault(overrides[CustomThemeKeys.Inactive], t.Inactive)
	t.Success = customColorOrDefault(overrides[CustomThemeKeys.Success], t.Success)
	t.Error = customColorOrDefault(overrides[CustomThemeKeys.Error], t.Error)
}

func customColorOrDefault(color string, def lipgloss.Color) lipgloss.Color {
	if color == "" {
		return def
	}

	return lipgloss.Color(color)
}
