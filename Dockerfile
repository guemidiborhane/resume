# Stage 1: Builder
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy source files
COPY . .

# Statically compile the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Final image
FROM scratch
LABEL org.opencontainers.image.source=https://github.com/guemidiborhane/resume

# Copy the statically compiled binary from the builder stage
COPY --from=builder /app/main /app/main

# Expose the port
EXPOSE 3000

# Add health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD ["/app/main", "healthcheck"]

# Set the entrypoint for the container
CMD ["/app/main"]
