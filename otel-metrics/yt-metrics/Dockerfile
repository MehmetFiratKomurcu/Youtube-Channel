# Build stage
FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

# Build the Go app
RUN go build -o yt-otel-metrics .

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the build stage
COPY --from=build /app/yt-otel-metrics .

EXPOSE 3001

CMD ["./yt-otel-metrics"]
