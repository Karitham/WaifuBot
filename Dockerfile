# Builder
FROM golang as builder
WORKDIR /build

# Setup gotip
RUN go install golang.org/dl/gotip@latest
RUN gotip download

COPY go.mod go.sum ./
RUN gotip mod download
COPY . .

RUN CGO_ENABLED=0 gotip build -ldflags="-w -s" -o /build/bot

# Runner
FROM alpine:3.15
COPY --from=builder /build/bot /bot
ENTRYPOINT ["/bot"]