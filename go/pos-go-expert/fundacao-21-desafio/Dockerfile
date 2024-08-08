# syntax=docker/dockerfile:1

FROM golang:1.22.5

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
RUN go mod tidy

# Copy the source code. Note the slash at the end, as explained in
COPY . .

#RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
#RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
#RUN apk install protobuf-compiler

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o des21 cmd/server/server.go

# Run
#ENTRYPOINT ["./des21"]
#CMD ["./des21"]



#FROM scratch
FROM alpine:latest
RUN mkdir /app

COPY --from=0 /app/des21 /app/des21

RUN apk --no-cache add ca-certificates

EXPOSE 8080
EXPOSE 50051
EXPOSE 5453

ENTRYPOINT ["/app/des21"]
#CMD ["/app/des21"]