FROM golang:1.22

WORKDIR /app

ADD . /app

RUN go get -d -v ./...

RUN go build -o main ./cmd

EXPOSE 8080

ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /app/wait-for-it.sh

RUN chmod +x /app/wait-for-it.sh

CMD ["/app/wait-for-it.sh", "db:${DB_PORT}", "--", "/app/main"]