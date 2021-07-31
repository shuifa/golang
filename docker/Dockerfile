FROM golang:latest

CMD mkdir /app

WORKDIR /app

ADD . /app

RUN go build -o main ./server.go

EXPOSE 8080

CMD /app/main