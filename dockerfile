FROM golang:1.19
WORKDIR /app
COPY . .
RUN go build -o main .
VOLUME /app/logs
CMD ["./main", "/app/logs/orders.txt"]