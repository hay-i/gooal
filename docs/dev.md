# Development Documentation

## Tech stack

Gooal is built on the "GOTH" stack. [Go](https://go.dev/) for the majority, [templ](https://templ.guide/) for our HTML & [HTMX](https://htmx.org/) for the interactivity.

Our CSS is using [SASS](https://sass-lang.com/) and are using very minimal JS to accomplish any front end functionality. At the time of writing, [sortable](https://github.com/SortableJS/Sortable) and HTMX are the only externl libraries we have minified in the codebase outside of any code we've written ourselves.

We've also chosen [MongoDB](https://www.mongodb.com/) for our database for it's flexible schemas.

## Installation

Providing you have the necessary development tools installed, you can get up and running using the `make` commands. Use `make help` to see a list of options.

1) Install Golang using [asdf](https://asdf-vm.com/guide/getting-started.html) / `asdf plugin add golang && asdf install`
2) Install Templ from the [docs](https://templ.guide/quick-start/installation) / `go install github.com/a-h/templ/cmd/templ@v0.2.598`
3) Install Docker
4) Install SASS CLI from the [docs](https://sass-lang.com/install/) / `brew install sass/sass/sass`

## Up and running

1) Start your docker.
2) Start the docker containers using `make up`.
3) Start the application with hot module reloading using `make hmr`.

Alternatively, if you don't need hot module reloading, you can just run `go run .` or `make start`

## Testing

To add tests, create a file with the suffix `_test.go` next to the file you're testing, and run `make test` to run the tests.

## Debugging

My current go-to debugging technique is to use our `pkg/logger` package. This will add some color (both literally and figuratively) to your logs when you're debugging. There is also an issue to upgrade from this makeshift debugger: https://github.com/hay-i/gooal/issues/67

## Pushing Code

Due to the way `templ`s watch feature works, the generated go code will look different when running a watch instead of a standard `templ generate`. The differently generated code would work in production, but it's not recommended to commit it as it's less performant.

Because of this, it's recommended to run `make build` before committing code. This script will compile your scss as well as generate the correct templ files to push up.

After writing your great commit messages, create a PR with closing tags for the issues you're solving and write a summary in the description.

### Workflows

To check the above (templ file generation & scss), we run workflows on each step per PR which have to pass before merging.

## Database

### Using MongoDB

To view a collection, enter the mongo cli with `make dbCli` and run the following:
```
use gooal
db.templates.find()
```

### [vim-dadbod](https://github.com/tpope/vim-dadbod)

To connect, add a connection to `mongodb://root:example@127.0.0.1:27017/gooal?authSource=admin`
