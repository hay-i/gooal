# chronologger

Chronologger is a personal review app, built with Golang.

# Development

Start the application with hot module reloading using `make hmr`.

Alternatively, if you don't need hot module reloading, you can just run `go run .` or `make start`

## Known issues

```
(!) templ version check failed: failed to parse go.mod file: ~/chronologger/go.mod:3: invalid go version '1.21.0': must match format 1.23
```
To fix this, make sure you have at least the following version installed:
```
go install github.com/a-h/templ/cmd/templ@v0.2.543
```
You can test this by running `templ version`.

# Dependencies
