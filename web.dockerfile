FROM golang:1.25.4-alpine3.22 AS Builder
WORKDIR /app
COPY . . 
RUN go mod download && go build -o server
CMD ["./server"]
