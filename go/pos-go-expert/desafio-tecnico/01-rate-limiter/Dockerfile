# syntax=docker/dockerfile:1

FROM golang:1.22.3

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code. Note the slash at the end, as explained in
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o rate-limiter cmd/server/main.go

# Optional:
EXPOSE 8080

# Run
#ENTRYPOINT ["./rate-limiter"]
CMD ["./rate-limiter"]