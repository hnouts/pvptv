FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download
# RUN apk update && apk add make \
#      rm -rf /var/cache/apk/*

# CMD make build
RUN GOARCH=wasm GOOS=js go build -o docs/web/app.wasm ./bin/arenatv
RUN go build -o docs/arenatv ./bin/arenatv
RUN chown -Rf "1000" ./*

FROM alpine

COPY --from=builder /app/docs /app
EXPOSE 8080
WORKDIR /app
CMD ["./arenatv"]
