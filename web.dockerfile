FROM golang:1.25.4-alpine3.22 AS builder 
# OS base,go compiler , go toolchain,etc etc 
WORKDIR /app
COPY . . 
RUN go mod download 
CMD ["go", "run","."]
