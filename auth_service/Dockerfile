FROM golang:1.20.5-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /build

FROM alpine:3.14

WORKDIR /
COPY --from=build /build /build

EXPOSE 8080

CMD [ "/build" ]