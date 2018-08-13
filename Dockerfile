FROM golang:1.10.3-alpine

EXPOSE 8000

RUN apk add --update git; \
    mkdir -p ${GOPATH}/go-rest-api; \
    go get -u github.com/gorilla/mux

WORKDIR ${GOPATH}/go-rest-api/
COPY src/rest-api/main.go ${GOPATH}/go-rest-api/
RUN go build -o go-rest-api .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=0 /go/go-rest-api/go-rest-api .

CMD [ "/app/go-rest-api" ]
