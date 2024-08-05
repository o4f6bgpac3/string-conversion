# Build stage
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

# Frontend build stage
FROM node:16-alpine AS frontend-builder
WORKDIR /app
COPY frontend .
RUN npm install && npm run build

# Final stage
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=frontend-builder /app/dist ./frontend/dist
COPY .env .
EXPOSE 8080
CMD ["./main", "serve"]