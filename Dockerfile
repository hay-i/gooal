FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go install github.com/a-h/templ/cmd/templ@v0.2.598

EXPOSE 8080

CMD ["go", "run", "/app/cmd/gooal/main.go"]
