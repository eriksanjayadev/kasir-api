# ===== BUILD STAGE =====
FROM golang:1.25.0-alpine AS builder
LABEL "language"="go"

WORKDIR /app

# copy module files first (better cache)
COPY go.mod go.sum ./
RUN go mod download

# copy source code
COPY . .

# build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

# ===== RUN STAGE =====
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server ./server

EXPOSE 8080

CMD ["./server"]
