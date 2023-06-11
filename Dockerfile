# Build the project
FROM golang:1.20 as builder

WORKDIR /go/src/github.com/qosimmax/file-server-api
ADD . .

RUN make build
RUN mkdir buckets

# Create production image for application with needed files
FROM golang:1.20.5-alpine3.18

RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/github.com/qosimmax/file-server-api .

CMD ["./bin/file-server-api"]
