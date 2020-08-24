# Build image builder
FROM golang:alpine as builder
WORKDIR /home/waifubot
COPY . .
RUN go build -i

# Build app image
FROM alpine:latest as image
WORKDIR /home/waifubot
COPY --from=builder /home/waifubot/WaifuBot* .
CMD [ "./WaifuBot" ]