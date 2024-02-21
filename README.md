# chronologger

Chronologger is a personal review app, built with Golang.

# Development

## Up and running

Start the application with hot module reloading using `make hmr`.

Alternatively, if you don't need hot module reloading, you can just run `go run .` or `make start`

## Using MongoDB

To view a collection, enter the mongo cli with `make dbCli` and run the following:
```
use chronologger
db.templates.find()
```
