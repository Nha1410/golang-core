FROM golang:1.23

WORKDIR /go/src

RUN go install github.com/air-verse/air@latest

ENV DEV_ENV=development

WORKDIR /app

CMD ["air"]
