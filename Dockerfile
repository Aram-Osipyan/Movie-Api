FROM golang:latest

WORKDIR api

# COPY go.mod Gopkg.lock ./

# RUN go mod download

COPY . .

RUN go build -o api .

EXPOSE 8080

CMD ["api"]