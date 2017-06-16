FROM golang:latest

WORKDIR /go/src/github.com/luke-davies/go-messages
COPY . .

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-messages"]
