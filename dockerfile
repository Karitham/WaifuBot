# Builder
FROM golang:alpine as builder
WORKDIR /build

# For Cache
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /build/bot

# Runner
FROM alpine:3.14
COPY --from=builder /build/bot /bin/bot
ENTRYPOINT ["/bin/bot"]
