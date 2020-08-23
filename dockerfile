FROM golang:alpine

# Set a directory for the app
WORKDIR /home/waifubot

# Copy all the files to the container
COPY . .

# build
RUN go build -i

# Expose port for mongo
EXPOSE 27017 

# Run the app
CMD [ "./WaifuBot" ]