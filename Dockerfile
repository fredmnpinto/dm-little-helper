# syntax=docker/dockerfile:1

FROM golang:alpine

COPY . ./app

WORKDIR ./app

RUN go get .

RUN go build -o /dm-little-helper

EXPOSE 80

CMD [ "/dm-little-helper" ]