# Stage 1 — The first stage uses the full Go image to compile your binary.
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/main.go

# Stage 2 — The second stage uses a tiny Alpine Linux image and copies only the compiled binary and your static files into it.
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/static ./static
COPY --from=builder /app/knowledge ./knowledge

EXPOSE 8080

CMD ["./server"]

# The result is a final image that's maybe 15–20MB instead of hundreds of MB. 
# Smaller images are faster to push, pull, and deploy — this is standard practice for production Go services.