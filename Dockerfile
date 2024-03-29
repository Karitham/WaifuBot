# Builder
FROM golang:1.19-alpine as builder

RUN apk add git

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /build/bot

# Runner
FROM alpine:3.15
COPY --from=builder /build/bot /bot
ENTRYPOINT ["/bot"]
