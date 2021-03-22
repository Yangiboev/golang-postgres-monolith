# workspace (GOPATH) configured at /go
FROM golang:1.13.1 as builder


#
RUN mkdir -p $GOPATH/src/github.com/Yangiboev/golang-postgres-monolith
WORKDIR $GOPATH/src/github.com/Yangiboev/golang-postgres-monolith

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    make build && \
    mv ./bin/api_gateway /



FROM alpine
COPY --from=builder api_gateway .
RUN mkdir config

ENTRYPOINT ["/api_gateway"]