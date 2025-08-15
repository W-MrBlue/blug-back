FROM golang:1.23.0-alpine AS builder
LABEL authors="mrblue"

ENV CGO_ENABLED=1
WORKDIR /app/build
COPY . .
RUN apk add  build-base
RUN go build -o blug main.go


FROM alpine:latest AS runner

WORKDIR /app/bin


COPY --from=builder /app/build .
EXPOSE 8080

CMD ["./blug"]