hmr: up sass ## Start the hot module replacement server
	templ generate --watch --proxy="http://localhost:1323" --cmd="go run ."

start: templ ## Generate templates and start the server
	go run .

templ: ## Generate the templates
	templ generate

up: ## Start the docker containers
	docker compose up -d

sass: ## Compile and minify SASS
	sass --watch --style compressed .

down: ## Stop the docker containers
	docker compose down

dbCli: ## Connect to the mongo db
	docker exec -it chronologger-mongo-1 mongosh --username root --password example

help:
	@sed -n -E "s/(^[^ ]+):.* ## (.*)/`printf "\033[32m"`\1|`printf "\033[0m"` \2/p" $(MAKEFILE_LIST) | sort | column -t -s '|'
