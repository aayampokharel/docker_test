FROM golang:1.25.4-alpine3.22 AS BUILDER
WORKDIR /app
COPY . .
CMD ["go","run","."]

