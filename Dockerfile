FROM golang:1.19 as builder
# Set necessary env variables needed for our image
ENV config=docker
ENV CGO_ENABLED=0

# Move to working directory /bin
WORKDIR /app

#  Copy and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy all folder and files to working directory
COPY . .

# Build
RUN go build -o app cmd/main.go

# Make dir for binary file
WORKDIR /bin

# Copy binary file from builder file
RUN cp /app/app .

# 2 stage build a smaller image
FROM alpine:3.16

RUN apk --no-cache add ca-certificates

COPY . .
COPY --from=builder /bin/app /

EXPOSE 5000

ENTRYPOINT ["/app"]