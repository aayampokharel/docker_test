FROM golang:1.26.0-alpine3.22 AS builder
WORKDIR /app
COPY . . 
RUN ["chmod","u+x","./cmd.sh"]
CMD ["./cmd.sh"]

