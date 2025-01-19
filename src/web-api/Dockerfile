FROM golang:1.23

WORKDIR /usr/local/app
COPY go.mod ./
COPY ./internal ./internal
RUN go mod download

COPY . .
RUN go build -v -o app .

CMD ["./app"]