FROM golang:1.18.2

WORKDIR /app

COPY . .

RUN go get -u github.com/go-sql-driver/mysql

RUN go build -o main .

CMD ["./main"]