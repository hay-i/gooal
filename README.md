<img src="assets/svg/logo.svg" width="128"/>

# Gooal
[![GO Test](https://github.com/hay-i/gooal/actions/workflows/go-test.yml/badge.svg)](https://github.com/hay-i/gooal/actions/workflows/go-test.yml/badge.svg)
[![GO Formatting](https://github.com/hay-i/gooal/actions/workflows/go-format.yml/badge.svg)](https://github.com/hay-i/gooal/actions/workflows/go-format.yml/badge.svg)
[![Templ Formatiing](https://github.com/hay-i/gooal/actions/workflows/templ-format.yml/badge.svg)](https://github.com/hay-i/gooal/actions/workflows/templ-format.yml/badge.svg)
[![Templ Compilation](https://github.com/hay-i/gooal/actions/workflows/templ-generate.yml/badge.svg)](https://github.com/hay-i/gooal/actions/workflows/templ-generate.yml/badge.svg)
[![SCSS Compilation](https://github.com/hay-i/gooal/actions/workflows/minify-css.yml/badge.svg)](https://github.com/hay-i/gooal/actions/workflows/minify-css.yml/badge.svg)

Gooal is a personal review app, built with Golang.

# Development

You can get up and running using `make`, use `make help` to see a list of options.

## Installation

1) Install Golang using [asdf](https://asdf-vm.com/guide/getting-started.html) / `asdf plugin add golang && asdf install`
2) Install Templ from the [docs](https://templ.guide/quick-start/installation) / `go install github.com/a-h/templ/cmd/templ@v0.2.598`
3) Install Docker
4) Install SASS CLI from the [docs](https://sass-lang.com/install/) / `brew install sass/sass/sass`

## Up and running

1) Start your docker.
2) Start the docker containers using `make up`.
3) Start the application with hot module reloading using `make hmr`.

Alternatively, if you don't need hot module reloading, you can just run `go run .` or `make start`

### Using MongoDB

To view a collection, enter the mongo cli with `make dbCli` and run the following:
```
use gooal
db.templates.find()
```

### Pushing Code

Due to the way `templ`s watch feature works, the generated go code will look different when running a watch instead of a standard `templ generate`.

Because of this, it's recommended to run `make build` before committing code.

Note: The differently generated code would work in production, but it's not recommended to commit it as it's less performant.

### Testing

To add tests, create a file with the suffix `_test.go` next to the file you're testing, and run `make test` to run the tests.
