FROM golang:1.22.4-bookworm

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app ./cmd/server/main.go

# Use a minimal base image for the final container
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /root/

# Copy the pre-built binary file from the builder stage
COPY --from=builder /app/my-go-app .

# Expose the port on which the app will run
EXPOSE 8080

# Command to run the executable
CMD ["./my-go-app"]

