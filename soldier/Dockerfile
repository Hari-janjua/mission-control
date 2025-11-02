FROM golang:1.24-alpine

WORKDIR /app

# Copy go modules
COPY go.mod go.sum ./../
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN go build -o /app/soldier-app ./soldier/cmd

# Run the binary
CMD ["/app/soldier-app"]
