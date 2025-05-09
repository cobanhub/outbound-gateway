# Start from official Go image
FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o madakaripura ./cmd/main.go

# Final image
FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=builder /app/madakaripura .

USER nonroot:nonroot

ENTRYPOINT ["./madakaripura"]
