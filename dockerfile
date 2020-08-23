FROM alpine

# Set a directory for the app
WORKDIR /home/waifubot

# Copy the binary to the container
COPY WaifuBot* .

# Expose port for mongo
EXPOSE 27017 

# Run the app
CMD [ "./WaifuBot" ]