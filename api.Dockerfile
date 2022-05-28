FROM golang:1.18.2

WORKDIR /app

COPY ./app .

RUN go get -u github.com/go-sql-driver/mysql

RUN go install github.com/githubnemo/CompileDaemon@latest 

ENTRYPOINT CompileDaemon --build='go build ./main.go' -command='./main'