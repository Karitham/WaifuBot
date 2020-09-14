# Build image builder
FROM golang:alpine as builder
WORKDIR /home/waifubot
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build

# Build app image
FROM alpine:latest as image
WORKDIR /home/waifubot
COPY --from=builder /home/waifubot/WaifuBot* .
CMD [ "./WaifuBot" ]