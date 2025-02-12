FROM golang:alpine as builder

WORKDIR /app

COPY . /app

RUN go build -o /app/bin/snippetbox /app

FROM alpine:latest as runner


WORKDIR /app 

# Copy binary
COPY --from=builder /app/bin /app

# Copy assets
COPY --from=builder /app/ui /app/ui


EXPOSE 8080

CMD ["/app/snippetbox", "--port", "8080"]