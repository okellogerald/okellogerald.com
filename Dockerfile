# Use the official Go image as the base image
FROM golang:1.20

RUN mkdir /build
WORKDIR /build

COPY . /build

RUN go get
RUN go build -o /docker-okellogerald.com

# Expose port 8080
EXPOSE 8080

# Define the entry point for the container
ENTRYPOINT [ "/docker-okellogerald.com" ]