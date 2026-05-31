FROM golang:1.26-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /build/prober ./cmd/prober

FROM scratch

COPY --from=builder /build/prober /prober
COPY --from=builder /build/config.yaml /config.yaml

ENTRYPOINT ["/prober"]
