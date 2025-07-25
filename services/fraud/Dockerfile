FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o fraud ./services/fraud/cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/fraud .
EXPOSE 50052
ENTRYPOINT ["./fraud"]