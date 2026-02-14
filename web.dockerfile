FROM golang:1.26.0-alpine3.22 AS builder
WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN ["go","mod","download"]
COPY . . 
RUN ["chmod","u+x","./cmd.sh"]
CMD ["./cmd.sh"]