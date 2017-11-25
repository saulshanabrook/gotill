# `gotill`

Runs a command until some output is reached, then exit, keeping the command running in the background.

## Install


```
go get github.com/saulshanabrook/gotill
```

TODO: Add binaries with [goreleaser](https://goreleaser.com/).

## Example

Run `jupyter lab --watch` in the foreground until the pattern `webpack --watch` is displayed, then send to background:

```bash
gotil 'webpack --watch' jupyter lab --watch
```
