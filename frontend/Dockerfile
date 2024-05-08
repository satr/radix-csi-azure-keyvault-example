FROM golang:1.22.3-alpine3.19 as builder

WORKDIR /app
COPY ./app /app
RUN apk update && apk upgrade
RUN apk add ca-certificates curl git
RUN apk add --no-cache gcc musl-dev

# Install project dependencies
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o /usr/local/bin/frontend
RUN addgroup -S -g 1700 radix-job-example
RUN adduser -S -u 1000 -G radix-job-example radix-job-example

USER 1000
EXPOSE 8082
ENTRYPOINT ["/usr/local/bin/frontend"]