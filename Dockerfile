FROM golang:1.13-alpine

WORKDIR /go/src/github.com/lekhacman/dms
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

LABEL version="0.0.1"
EXPOSE 8001

CMD ["dms"]
