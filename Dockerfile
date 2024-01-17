FROM golang:1.18-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go build -o project-cuba

CMD ["./project-cuba"]