FROM golang:1.24-alpine

WORKDIR /app

RUN apk update && apk upgrade
RUN apk add sqlite

RUN go mod donwload
RUN go build -v -o server

CMD [ "./server" ]