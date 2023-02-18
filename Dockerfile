ARG GO_VERSION=1.18.3

FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./

RUN go build -o main cmd/main.go

RUN chmod +x main

FROM alpine AS final

WORKDIR /app

COPY --from=builder app/main main

ENTRYPOINT ["./main"]