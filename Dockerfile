# syntax=docker/dockerfile:1

FROM golang:alpine

COPY . ./app

WORKDIR ./app

RUN go get .

RUN go build -o /dm-little-helper

EXPOSE 8080

CMD [ "/dm-little-helper" ]