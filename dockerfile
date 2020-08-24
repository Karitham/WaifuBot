# Build image builder
FROM golang:alpine as builder
LABEL maintainer="PL Pery <plouis.pery@gmail.com>"
WORKDIR /home/waifubot
COPY . .
RUN go build -i

# Build app image
FROM alpine:latest as image
WORKDIR /root/
COPY --from=builder /home/waifubot/WaifuBot* .
CMD [ "./WaifuBot" ]