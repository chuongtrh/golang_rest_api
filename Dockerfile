# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/libs/dockerize /usr/local/bin
COPY --from=builder /app/src/config/auth_model.conf ./src/config/auth_model.conf
COPY --from=builder /app/src/config/auth_policy.csv ./src/config/auth_policy.csv

# Build Args
#ARG LOG_DIR=./logs

# Create Log Directory
#RUN mkdir -p ${LOG_DIR}

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD dockerize -wait tcp://mydb:5432 -timeout 120s ./main
#CMD ["./main"]