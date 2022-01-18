FROM golang:1.17.6-alpine3.15 as builder

RUN apk update && \
    apk add ca-certificates curl git  && \
    apk add --no-cache gcc musl-dev && \
    go get -u golang.org/x/lint/golint github.com/frapposelli/wwhrd

WORKDIR ..

# Install project dependencies
COPY ./app ./app
WORKDIR ./app
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o /usr/local/bin/frontend
RUN addgroup -S -g 1700 radix-job-example
RUN adduser -S -u 1000 -G radix-job-example radix-job-example

USER 1000
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/frontend"]