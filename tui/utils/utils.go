package utils

import (
	"github.com/alecthomas/chroma/styles"
	"github.com/charmbracelet/lipgloss"
)

var (
	Vim          = styles.Vim.Name
	ParaisoLight = styles.ParaisoLight.Name
	ParaisoDark  = styles.ParaisoDark.Name
	MonakiDark   = styles.Monokai.Name
	MonkaiLight  = styles.MonokaiLight.Name
)

func GetChromaTheme() string {

	if lipgloss.HasDarkBackground() {
		return Vim
	}
	return ParaisoLight
}
