# Start from the latest golang base image
FROM golang:1.20.7-alpine3.18 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY go.mod .
COPY /src .
COPY /profiles .
COPY .env .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api_app ./src/cmd/api/*.go

######## Start a new stage from scratch #######
FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/api_app .
COPY --from=builder /app/profiles ./profiles
COPY --from=builder /app/.env .

# Expose port 8083 to the outside world
EXPOSE 8083

# Command to run the executable
CMD ["./api_app"]
