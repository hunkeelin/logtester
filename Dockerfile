FROM alpine:latest
WORKDIR /app
COPY . .
CMD ["./run"]
