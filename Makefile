hmr: ## Start the hot module replacement server
	$(MAKE) -j2 sassWatch templProxy

templProxy: ## Start the hot module replacement server for templ
	templ generate --watch --proxy="http://localhost:1323" --cmd="go run ."

sassWatch: ## Compile, minify and watch SASS changes
	sass --style compressed --watch .

sassGen: ## Compile and minify SASS
	sass --style compressed .

templ: ## Generate the templates
	templ generate

build: templ sassGen ## Generate the templates and compile the SASS

start: ## Generate templates and start the server
	go run .

up: ## Start the docker containers
	docker compose up -d

down: ## Stop the docker containers
	docker compose down

dbCli: ## Connect to the mongo db
	docker exec -it chronologger-mongo-1 mongosh --username root --password example

help:
	@sed -n -E "s/(^[^ ]+):.* ## (.*)/`printf "\033[32m"`\1|`printf "\033[0m"` \2/p" $(MAKEFILE_LIST) | sort | column -t -s '|'
