FROM golang:1.23-alpine AS builder

WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/server ./cmd/server

FROM alpine:3.20

RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/configs/config.yaml ./configs/

EXPOSE 8080

CMD ["./server", "-config", "configs/config.yaml"]
