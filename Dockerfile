FROM golang:1.11.5-stretch

WORKDIR /go/src/api
COPY . .

RUN curl https://glide.sh/get | sh

RUN glide install

CMD cd api && go run main.go
