FROM golang:1.20.5-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN chmod +x build

RUN go build -o ./build ./cmd/main.go

USER root

EXPOSE 8080

CMD [ "./build/main" ]