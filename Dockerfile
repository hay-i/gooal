FROM golang:1.22-alpine

RUN apk add --no-cache make 

WORKDIR /app

COPY . .

RUN go mod download

RUN go install github.com/a-h/templ/cmd/templ@v0.2.598

EXPOSE 8080

CMD ["make", "start"]
