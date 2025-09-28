# jqp

a TUI playground for exploring jq.

![jqp](https://user-images.githubusercontent.com/23270779/191256434-05aeda9d-9ee2-4b5e-a23f-6548dac08fdb.gif)

This application utilizes [itchyny's](https://github.com/itchyny) implementation of `jq` written in Go, [`gojq`](https://github.com/itchyny/gojq).

## Installation

### homebrew

```bash
brew install jqp
```

### macports

```bash
sudo port install jqp
```

### Arch Linux
Available through the Arch User Repository as [jqp-bin](https://aur.archlinux.org/packages/jqp-bin).
```bash
yay -S jqp-bin
```

### Snap install
<a href="https://snapcraft.io/jqp"><img src="https://snapcraft.io/jqp/badge.svg" alt="Snap Status"></a>

```
sudo snap install jqp
```

### GitHub releases

Download the relevant asset for your operating system from the latest GitHub release. Unpack it, then move the binary to somewhere accessible in your `PATH`, e.g. `mv ./jqp /usr/local/bin`.

### Build from source

Clone this repository, build from source with `cd jqp && go build`, then move the binary to somewhere accessible in your `PATH`, e.g. `mv ./jqp /usr/local/bin`.

## Usage

```
➜ jqp --help
jqp is a terminal user interface (TUI) for exploring the jq command line utility.

You can use it to run jq queries interactively. If no query is provided, the interface will prompt you for one.

The command accepts an optional query argument which will be executed against the input JSON or newline-delimited JSON (NDJSON).
You can provide the input JSON or NDJSON either through a file or via standard input (stdin).

Usage:
  jqp [query] [flags]

Flags:
      --config string   path to config file (default is $HOME/.jqp.yaml)
  -f, --file string     path to the input JSON file
  -h, --help            help for jqp
  -t, --theme string    jqp theme
  -v, --version         version for jqp
```

`jqp` also supports input from STDIN. STDIN takes precedence over the command-line flag. Additionally, you can pass an optional query argument to jqp that it will execute upon loading.

```
➜ curl "https://api.github.com/repos/jqlang/jq/issues" | jqp '.[] | {"title": .title, "url": .url}'
```

> [!NOTE]
> Valid JSON or NDJSON [(newline-delimited JSON)](https://jsonlines.org/) can be provided as input to `jqp`.

## Keybindings

| **Keybinding** | **Action** |
|:---------------|:-----------|
| `tab` | cycle through sections |
| `shift-tab` | cycle through sections in reverse |
| `ctrl-y` | copy query to system clipboard[^1] |
| `ctrl-s` | save output to file (copy to clipboard if file not specified) |
| `ctrl-t` | toggle showing/hiding input panel |
| `ctrl-c` | quit program / kill long-running query |

### Query Mode

| **Keybinding** | **Action** |
|:---------------|:-----------|
| `enter` | execute query |
| `↑`/`↓` | cycle through query history |
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

`jqp` can be configured with a configuration file. By default, `jqp` will search your home directory for a YAML file named `.jqp.yaml`. A path to a YAML configuration file can also be provided to the `--config` command-line flag.

```bash
➜ jqp --config ~/my_jqp_config.yaml < data.json
```

If a configuration option is present in both the configuration file and the command-line, the command-line option takes precedence. For example, if a theme is specified in the configuration file and via `-t/--theme flag`, the command-line flag will take precedence.

### Available Configuration Options

```yaml
theme:
  name: "nord" # controls the color scheme
  chromaStyleOverrides: # override parts of the chroma style
    kc: "#009900 underline" # keys use the chroma short names
```

## Themes

Themes can be specified on the command-line via the `-t/--theme <themeName>` flag. You can also set a theme in your [configuration file](#configuration).

```yaml
theme:
  name: "monokai"
```

<img width="1624" alt="Screen Shot 2022-10-02 at 5 31 40 PM" src="https://user-images.githubusercontent.com/23270779/193477383-db5ca769-12bf-4fd0-b826-b1fd4086eac3.png">

### Chroma Style Overrides

Overrides to the chroma styles used for a theme can be configured in your [configuration file](#configuration).

For the list of short keys, see [`chroma.StandardTypes`](https://github.com/alecthomas/chroma/blob/d38b87110b078027006bc34aa27a065fa22295a1/types.go#L210-L308). To see which token to use for a value, see the [JSON lexer](https://github.com/alecthomas/chroma/blob/master/lexers/embedded/json.xml) (look for `<token>` tags). To see the color and what's used in the style you're using, look for your style in the chroma [styles directory](https://github.com/alecthomas/chroma/tree/master/styles).

```yaml
theme:
  name: "monokai" # name is required to know which theme to override
  chromaStyleOverrides:
    kc: "#009900 underline"
```

You can change non-syntax colors using the `styleOverrides` key:
```yaml
theme:
  styleOverrides:
    primary: "#c4b28a"
    secondary: "#8992a7"
    error: "#c4746e"
    inactive: "#a6a69c"
    success: "#87a987"
```

Themes are broken up into [light](#light-themes) and [dark](#dark-themes) themes. Light themes work best in terminals with a light background and dark themes work best in a terminal with a dark background. If no theme is specified or a non-existent theme is provided, the default theme is used, which was created to work with both terminals with a light and dark background.

### Light Themes

- `abap`
- `algol`
- `arduino`
- `autumn`
- `borland`
- `catppuccin-latte`
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
- `paraiso-light`
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
- `catppuccin-frappe`
- `catppuccin-macchiato`
- `catppuccin-mocha`
- `doom-one`
- `doom-one2`
- `dracula`
- `fruity`
- `github-dark`
- `gruvbox`
- `monokai`
- `native`
- `paraiso-dark`
- `rrt`
- `solarized-dark`
- `solarized-dark256`
- `swapoff`
- `vim`
- `witchhazel`
- `xcode-dark`

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
