FROM golang:1.18beta2-alpine as builder

RUN apk add git

WORKDIR /build

# For Cache
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /build/bot

# Runner
FROM alpine:3.15
COPY --from=builder /build/bot /bin/bot
ENTRYPOINT ["/bin/bot"]