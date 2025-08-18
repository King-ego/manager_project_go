FROM golang:1.24.4

# Set the working directory
WORKDIR /src/app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the source code
COPY . .

# Expose the port the app runs on
EXPOSE 8989

# Build the Go app
RUN go build -o main cmd/main.go

RUN go build -o migrate cmd/migrate/main.go
RUN go build -o rollback cmd/rollback/main.go

# Command to run the executable
CMD ["./main"]