FROM golang:1.23

WORKDIR /usr/local/app
COPY go.mod go.sum ./
COPY ./internal ./internal
RUN go mod download

COPY . .
RUN go build -v -o app .

CMD ["./app"]