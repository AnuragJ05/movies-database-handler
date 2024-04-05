# use official Golang image
FROM golang:1.16.3-alpine3.13

# set working directory
WORKDIR /workspace

# Copy the source code
COPY . . 

# Download and install the dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build .

#EXPOSE the port
EXPOSE 8000

# Run the executable
CMD ["./database-handler"]
