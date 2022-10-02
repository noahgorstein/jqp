# jqp 

a TUI playground for exploring jq.

![jqp](https://user-images.githubusercontent.com/23270779/191256434-05aeda9d-9ee2-4b5e-a23f-6548dac08fdb.gif)

This application utilizes [itchyny's](https://github.com/itchyny) implementation of `jq` written in Go, [`gojq`](https://github.com/itchyny/gojq).

## Installation

### homebrew

```bash
brew install noahgorstein/tap/jqp
```

### macports

```bash
sudo port install jqp
```

### Github releases

Download the relevant asset for your operating system from the latest Github release. Unpack it, then move the binary to somewhere accessible in your `PATH`, e.g. `mv ./jqp /usr/local/bin`.

### Build from source

Clone this repo, build from source with `cd jqp && go build`, then move the binary to somewhere accessible in your `PATH`, e.g. `mv ./jqp /usr/local/bin`.

## Usage

```
➜ jqp --help
jqp is a TUI to explore the jq command line utility

Usage:
  jqp [flags]

Flags:
      --config string   config file (default is $HOME/.jqp.yaml)
  -f, --file string     path to the input JSON file
  -h, --help            help for jqp
  -t, --theme string    jqp theme
  -v, --version         version for jqp
```

`jqp` also supports input from STDIN. 

```
➜ curl "https://api.github.com/repos/stedolan/jq/issues?per_page=2" | jqp 
```

## Keybindings 

| **Keybinding** | **Action** |
|:---------------|:-----------|
| `tab` | cycle through sections |
| `shift-tab` | cycle through sections in reverse |
| `ctrl-y` | copy query to system clipboard[^1] |
| `ctrl-s` | save output to file |
| `ctrl-c` | quit program / kill long running query |

### Query Mode

| **Keybinding** | **Action** |
|:---------------|:-----------|
| `enter` | execute query |
| `ctrl-a` | go to beginning of line |
| `ctrl-e` | go to end of line |
| `←`/`ctrl-b` | move cursor one character to left |
| `→`/`ctrl-f`| move cursor one character to right |
| `ctrl-k` | delete text after cursor line |
| `ctrl-u` | delete text before cursor |
| `ctrl-w` | delete word to left |
| `ctrl-d` | delete character under cursor |

### Input Preview and Output Mode

| **Keybinding** | **Action** |
|:---------------|:-----------|
| `↑/k` | up |
| `↓/j` | down |
| `ctrl-u` | page up |
| `ctrl-d` | page down |

## Configuration

`jqp` can be configured with a configuration file. By default, `jqp` will search your home directory for a YAML file named `.jqp.yaml`. A path to a YAML configuration file can also be provided to the `--config` command line flag. 

```bash
➜ jqp --config ~/my_jqp_config.yaml < data.json
```

If a configuration option is present in both the configuration file used and on the command line, the command line option takes precedence. For example, if a theme is specified in the configuration file and via `-t/--theme flag`, the command line flag will take precedence.

### Available Configuration

```yaml
theme: "nord" # controls the color scheme
file: "/path/to/input/file.json" # stdin takes precedence over command line flag and this option
```

## Themes

Themes can be specified on the command line via the `-t/--theme <themeName>` flag. You can also set a theme in your [configuration file](#configuration). 

```yaml
theme: "monakai"
```

Themes are broken up into "light" and "dark" themes below. Light themes work best in terminals with a light background and dark themes work best in a terminal with a dark background. If no theme is specified or non-existant theme is provided, the default theme is used, which was created to work with both terminals with a light and dark background. 

### Light Themes

- `abap`
- `algol`
- `arduino`
- `autumn`
- `borland`
- `colorful`
- `emacs`
- `friendly`
- `github`
- `gruvbox-light`
- `hrdark`
- `igor`
- `lovelace`
- `manni`
- `monokai-light`
- `murphy`
- `onesenterprise`
- `paradaiso-light`
- `pastie`
- `perldoc`
- `pygments`
- `solarized-light`
- `tango`
- `trac`
- `visual_studio`
- `vulcan`
- `xcode`

### Dark Themes

- `average`
- `base16snazzy`
- `doom-one` 
- `doom-one2` 
- `dracula`
- `fruity`
- `github-dark`
- `gruvbox`
- `monokai`
- `native`
- `paradaiso-dark`
- `rrt`
- `solarized-dark`
- `solarized-dark256`
- `swapoff`
- `vim`
- `witchhazel`

## Built with:

- [Bubbletea](https://github.com/charmbracelet/bubbletea)
- [Bubbles](https://github.com/charmbracelet/bubbles)
- [Lipgloss](https://github.com/charmbracelet/lipgloss)
- [gojq](https://github.com/itchyny/gojq)
- [chroma](https://github.com/alecthomas/chroma)

## Credits

- [jqq](https://github.com/jcsalterego/jqq) for inspiration

--------

[^1]: `jqp` uses [https://github.com/atotto/clipboard](https://github.com/atotto/clipboard) for clipboard functionality. Things should work as expected with OSX and Windows. Linux, Unix require `xclip` or `xsel` to be installed.
