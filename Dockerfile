FROM golang:1.24.4

WORKDIR /src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8989

RUN go build -o main cmd/main.go

RUN go build -o migrate cmd/migrate/main.go
RUN go build -o rollback cmd/rollback/main.go

CMD ["./main"]