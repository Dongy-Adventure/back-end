FROM golang:1.22.3

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY init.js /docker-entrypoint-initdb.d/

COPY . .

RUN go build -o server .

EXPOSE 3001

CMD ["./server"]
