FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o regretbox .

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/

COPY --from=builder /app/regretbox .
COPY --from=builder /app/templates ./templates/
COPY --from=builder /app/static ./static/

EXPOSE 8080
CMD ["./regretbox"]
