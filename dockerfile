FROM alpine:latest

WORKDIR /home/waifubot

COPY WaifuBot* .

CMD [ "./WaifuBot" ]