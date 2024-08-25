FROM golang:1.22.4-alpine as builder
RUN apk add --no-cache gcc musl-dev sqlite-dev
WORKDIR /app
COPY . .
RUN CGO_ENABLED=1 go build -ldflags="-w -s" -o server ./cmd/app/main.go 

FROM alpine:latest
RUN apk add --no-cache sqlite-libs
COPY --from=builder /app/server /server
COPY --from=builder /app/sql /sql
ENV PORT=8080
CMD ["./server"]