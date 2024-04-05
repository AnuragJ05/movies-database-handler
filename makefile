all: test build

test:
	go test ./...

fmt:
	go fmt ./...

build: 
	go build .

# Build the docker image
docker-build: 
	docker build . -t ${IMG}

# Push the docker image
docker-push:
	docker push ${IMG}