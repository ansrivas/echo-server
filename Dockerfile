    # Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app
# Copy go mod and sum files
COPY go.mod go.sum ./
# Download dependencies
RUN go mod download &&\
    apk update && apk add --no-cache upx
# Copy the source code
COPY *.go ./
# Build a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o echo-server . &&\
    upx echo-server

# Final stage
FROM gcr.io/distroless/static:nonroot
WORKDIR /
# Copy the binary from the builder stage
COPY --from=builder /app/echo-server /echo-server
# Run the binary
ENTRYPOINT ["/echo-server"]