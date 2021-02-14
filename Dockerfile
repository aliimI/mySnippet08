FROM golang

ADD . /go/src/

WORKDIR /app

COPY . /app

RUN go build -o app cmd/web/*

CMD ["/app/app"]