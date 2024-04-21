# Use a smaller base image for the builder stage
FROM golang:1.22.1 AS builder

# Set working directory
WORKDIR /app

# Download Go modules
COPY ./go.mod ./go.sum ./
RUN go mod download

# Copy the source code .dockerignore ignores senstitve files
COPY . .

# Build go application
RUN CGO_ENABLED=0 GOOS=linux go build -o wallmarkt-product-service ./cmd

# Create nonroot user
FROM alpine:latest AS security_provider
RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot

# Use a minimal base image for the final image
FROM scratch AS production
# Copy the nonroot user
COPY --from=security_provider /etc/passwd /etc/passwd
# Copy the built binary from the builder stage
COPY --from=builder /app/wallmarkt-product-service /wallmarkt-product-service

# Set nonroot user
USER nonroot

# Expose the port
EXPOSE 8081

# Run the binary
CMD ["/wallmarkt-product-service"]