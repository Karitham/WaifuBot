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
FROM scratch

WORKDIR /app

# Copy our static executable.
COPY --from=builder /app/app ./app

# Expose port
EXPOSE $API_PORT

# Run the binary
ENTRYPOINT ["/app/app"]