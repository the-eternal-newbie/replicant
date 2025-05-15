# Dockerfile

# Stage 1: Build the Go binary
FROM golang:1.24.3 as builder

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the source code
COPY . ./

# Build the Go CLI binary
RUN go build -o /replicant

# Stage 2: Slim runtime container
FROM debian:bookworm

WORKDIR /

# Copy compiled binary
COPY --from=builder /replicant /replicant

# Copy any data assets (e.g., BoltDB file or folder structure)
COPY data /data


VOLUME ["/data"]

CMD ["/replicant"]
