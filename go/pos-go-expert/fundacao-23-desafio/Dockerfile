# syntax=docker/dockerfile:1

FROM golang:1.22.5

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
RUN go mod tidy

# Copy the source code. Note the slash at the end, as explained in
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o weather cmd/main.go

# Run
#ENTRYPOINT ["./weather"]
#CMD ["./weather"]



#FROM scratch
FROM alpine:latest
RUN mkdir /app

COPY --from=0 /app/weather /app/weather

RUN apk --no-cache add ca-certificates

EXPOSE 8080

#ENTRYPOINT ["/app/weather"]
CMD ["/app/weather"]