FROM golang:1.24-alpine

WORKDIR /app

RUN apk update && apk upgrade
RUN apk add sqlite

CMD [ "./server" ]