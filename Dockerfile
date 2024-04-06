# Use official Golang image for building the Go application
FROM golang:latest

WORKDIR /app

COPY . .

# Remove existing go.mod and go.sum
RUN rm go.mod go.sum

# Initialize and tidy Go modules
RUN go mod init movies-database-handler && go mod tidy

# Build the Go app
RUN go build -o movies-database-handler .

RUN mkdir /tmp/astra/

EXPOSE 5000

CMD ["./movies-database-handler"]
