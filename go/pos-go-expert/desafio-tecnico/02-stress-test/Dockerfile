# syntax=docker/dockerfile:1

FROM golang:1.22.2

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code. Note the slash at the end, as explained in
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o stress-test cmd/cli/main.go

# Run
#ENTRYPOINT ["./rate-limiter"]
CMD ["./stress-test"]