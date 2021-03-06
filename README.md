## GoTube

> :metro: A terminal app for checking TFL tube status and last trains, written in Go

```console
// Show the status of all tube lines
$ gotube
```

<img src="screenshots/status.png"/>

```console
// Show the last trains for a particular station
$ gotube kings cross
```

<img src="screenshots/search.png"/>

### Install

To install, use `go get`:

```console
$ go get -d github.com/bluedaniel/gotube
```

### Build

To build a compressed executable, install upx -  `brew install upx`

```console
$ go build -ldflags "-s -w"
$ upx ./gotube
```


- - -

> Note: The terminal in the screenshot is made up of:
- Shell: [ZSH - BobbyRussel theme](https://github.com/robbyrussell/oh-my-zsh)
- Color scheme: [Dracula](https://draculatheme.com/iterm)
- Font: [Source Code Pro](https://github.com/adobe-fonts/source-code-pro)
