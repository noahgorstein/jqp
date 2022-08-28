package cmd

import (
	"errors"
	"os"

	"github.com/charmbracelet/bubbletea"
	"github.com/noahgorstein/jqp/tui/bubbles/jqplayground"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Version:      "0.1",
	Use:          "jqp",
	Short:        "jqp is a TUI to explore jq",
	Long:         `jqp is a TUI to explore the jq command line utility`,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		if isInputFromPipe() {
			stdin := streamToBytes(os.Stdin)

			isValidJson := isValidJson(stdin)
			if !isValidJson {
				return errors.New("JSON is not valid")
			}

			bubble := jqplayground.New(stdin, "STDIN")
			p := tea.NewProgram(bubble, tea.WithAltScreen())
			if err := p.Start(); err != nil {
				return err
			}
			return nil
		} else {

			// get the file
			file, e := getFile()
			if e != nil {
				return e
			}
			defer file.Close()

			// read the file
			data, err := os.ReadFile(flags.filepath)
			if err != nil {
				return err
			}

			isValidJson := isValidJson(data)
			if !isValidJson {
				return errors.New("JSON is not valid")
			}

			// get file info so we can get the filename
			fi, err := os.Stat(flags.filepath)
			if err != nil {
				return err
			}

			bubble := jqplayground.New(data, fi.Name())
			p := tea.NewProgram(bubble, tea.WithAltScreen())

			if err := p.Start(); err != nil {
				return err
			}
			return nil
		}

	},
}

var flags struct {
	filepath string
}

var flagsName = struct {
	file, fileShort string
}{
	"file", "f",
}

func Execute() {
	rootCmd.Flags().StringVarP(
		&flags.filepath,
		flagsName.file,
		flagsName.fileShort,
		"", "path to the input JSON file")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
