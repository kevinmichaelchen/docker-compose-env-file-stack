# syntax=docker/dockerfile:1

FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o ./app .

FROM alpine:3.14
COPY --from=builder /app/app .
USER 65534
ENTRYPOINT [ "./app" ]