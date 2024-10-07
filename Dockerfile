# Stage 1
FROM golang:1.23-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o cc-service2 .

# Stage 2
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/cc-service2 /app/.env ./

CMD ["./cc-service2"]