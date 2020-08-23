FROM alpine:latest

WORKDIR /home/waifubot

COPY WaifuBot* .

EXPOSE 27017 

CMD [ "./WaifuBot" ]