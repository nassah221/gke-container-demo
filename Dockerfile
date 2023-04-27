FROM golang:1.20.3-alpine

RUN mkdir app
WORKDIR /app
COPY . .

RUN go build -o app
CMD ./app
