FROM golang:alpine

ADD . /app
WORKDIR /app
RUN go build -o out ./app

CMD ["/app/out"]