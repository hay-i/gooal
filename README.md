# chronologger

Chronologger is a personal review app, built using Golang.

# Development

If you have `make` you can run the app with hot module reloading using `make hmr`.

If you don't need hot module reloading, you can just run `go run .`

## Known issues

```
(!) templ version check failed: failed to parse go.mod file: ~/chronologger/go.mod:3: invalid go version '1.21.0': must match format 1.23
```
To fix this, make sure you have the following version enabled:
```
go install github.com/a-h/templ/cmd/templ@v0.2.543
```
You can test this by running `templ version`. I had to setup my `GOPATH` locally before I installed it. Not sure why, since go modules should stop this being needed.

# Dependencies

