FROM golang:1.18-alpine3.17 as builder
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"'

FROM alpine:3.17
COPY --from=builder /go/bin/near-exporter /near-exporter
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/near-exporter"]
