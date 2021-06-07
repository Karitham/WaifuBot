# Build image builder
FROM golang:alpine as builder

# Set workdir
WORKDIR /app

# Copy folder content
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the remaining files
COPY . .

# Build the binary
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /app/app 

# Build app image.
# This image doesn't have a shell
FROM alpine:3.13

WORKDIR /app

# Copy our static executable.
COPY --from=builder /app/app ./app

# Run the binary
ENTRYPOINT ["/app/app"]