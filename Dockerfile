FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Andreas Leicher <andreas@kotaico.de>"

# install go-bindata
RUN go get -u github.com/go-bindata/go-bindata/...

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

#generate the schema
RUN go generate ./...

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/graphql ./graphql

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/bin/graphql .

ENV SQS_QUEUE_NAME go-graphql-sqs-example
# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./graphql"]
