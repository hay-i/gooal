# chronologger

Chronologger is a personal review app, built with Golang.

# Development

You can get up and running using `make`, use `make help` to see a list of options.

## Installation

1) Install Golang using ASDF
2) Install Templ from [docs](https://templ.guide/quick-start/installation)
3) Install Docker

## Up and running

Start the application with hot module reloading using `make hmr`.

Alternatively, if you don't need hot module reloading, you can just run `go run .` or `make start`

## Using MongoDB

To view a collection, enter the mongo cli with `make dbCli` and run the following:
```
use chronologger
db.templates.find()
```
