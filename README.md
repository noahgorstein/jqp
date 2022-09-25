# jqp 

a TUI playground for exploring jq.

![jqp](https://user-images.githubusercontent.com/23270779/191256434-05aeda9d-9ee2-4b5e-a23f-6548dac08fdb.gif)

This application utilizes [itchny's](https://github.com/itchyny) implementation of `jq` written in Go, [`gojq`](https://github.com/itchyny/gojq).

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
  -f, --file string   path to the input JSON file
  -h, --help          help for jqp
  -v, --version       version for jqp
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
| `ctrl-c` | quit program |

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
