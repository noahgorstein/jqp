package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
	"github.com/noahgorstein/jqp/tui/bubbles/jqplayground"
	"github.com/noahgorstein/jqp/tui/theme"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Version:      "0.2.0",
	Use:          "jqp",
	Short:        "jqp is a TUI to explore jq",
	Long:         `jqp is a TUI to explore the jq command line utility`,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			// Apply the viper config value to the flag when the flag is not set and viper has a value
			if !f.Changed && viper.IsSet(f.Name) {
				val := viper.Get(f.Name)
				cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
			}
		})

		if isInputFromPipe() {
			stdin := streamToBytes(os.Stdin)

			isValidJson := isValidJson(stdin)
			if !isValidJson {
				return errors.New("JSON is not valid")
			}

			bubble := jqplayground.New(stdin, "STDIN", theme.GetTheme(flags.theme))
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

			bubble := jqplayground.New(data, fi.Name(), theme.GetTheme(flags.theme))
			p := tea.NewProgram(bubble, tea.WithAltScreen())

			if err := p.Start(); err != nil {
				return err
			}
			return nil
		}

	},
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory
		viper.AddConfigPath(home)

    // register the config file
		viper.SetConfigName(".jqp")

    //only read from yaml files
		viper.SetConfigType("yaml")

	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Config file:", viper.ConfigFileUsed(), "was used.")
	}
}

var flags struct {
	filepath string
	theme    string
}

var flagsName = struct {
	file       string
	fileShort  string
	theme      string
	themeShort string
}{
	file:       "file",
	fileShort:  "f",
	theme:      "theme",
	themeShort: "t",
}

var cfgFile string

func Execute() {

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "path to config file (default is $HOME/.jqp.yaml)")

	rootCmd.Flags().StringVarP(
		&flags.filepath,
		flagsName.file,
		flagsName.fileShort,
		"", "path to the input JSON file")

	rootCmd.Flags().StringVarP(
		&flags.theme,
		flagsName.theme,
		flagsName.themeShort,
		"", "jqp theme")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
