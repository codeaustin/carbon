FROM golang:1.8

WORKDIR /go/src/github.com/codeaustin/carbon
COPY . .

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]
EXPOSE 3000
