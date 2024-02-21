hmr: ## Start the hot module replacement server
	templ generate --watch --proxy="http://localhost:1323" --cmd="go run ."

start: ## Start the server
	go run .

up: ## Start the docker containers
	docker compose up -d

down: ## Stop the docker containers
	docker compose down

dbCli: ## Connect to the mongo db
	docker exec -it chronologger-mongo-1 mongosh --username root --password example

help:
	@sed -n -E "s/(^[^ ]+):.* ## (.*)/`printf "\033[32m"`\1|`printf "\033[0m"` \2/p" $(MAKEFILE_LIST) | sort | column -t -s '|'
