# # Setting a base image 
FROM golang:1.19.13-alpine3.18

Run addgroup app && adduser -S -G app app
USER app

# # Set the workdir and copy the files (from the workdir) to the docker image
WORKDIR /app
COPY . .

# # Building the source files
RUN go build -o app

#Define my port
EXPOSE 8080

CMD ["./app"]