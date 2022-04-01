# syntax=docker/dockerfile:1

FROM golang:1.17.8

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o ./main .

EXPOSE 8080

CMD [ "./main" ]