FROM golang:latest

WORKDIR /app
COPY . .

RUN go mod tidy

EXPOSE 8888
CMD go run .