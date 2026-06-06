
FROM golang:1.26-alpine AS builder

WORKDIR /build

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o sentry ./cmd/sentry

FROM scratch
COPY --from=builder /build/sentry /sentry
ENTRYPOINT ["/sentry"]
EXPOSE 8080