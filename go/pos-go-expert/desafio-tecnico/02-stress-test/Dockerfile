# syntax=docker/dockerfile:1

FROM golang:1.22.2

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
RUN go mod tidy

# Copy the source code. Note the slash at the end, as explained in
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o stress-test cmd/cli/main.go

# Run
#ENTRYPOINT ["./stress-test"]
#CMD ["./stress-test"]

FROM scratch
COPY --from=0 /app/stress-test /bin/stress-test
ENTRYPOINT ["/bin/stress-test"]
#CMD ["stress-test"]