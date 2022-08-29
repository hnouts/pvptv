FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN apk update && apk add make \
     rm -rf /var/cache/apk/*

CMD make run

EXPOSE 8080