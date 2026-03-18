FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o openmock ./cmd/openmock

FROM scratch

COPY --from=builder /app/openmock /openmock

EXPOSE 8080

ENTRYPOINT ["/openmock"]
