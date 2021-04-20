FROM golang:1.15.7-buster

RUN go get github.com/gorilla/mux

RUN mkdir /app
RUN cd src
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute
## any further commands inside our /app/src
## directory
WORKDIR /app/src

RUN go mod download
## we run go build to compile the binary
## executable of our Go program
RUN go build -o main .
## Our start command which kicks off
## our newly created binary executable
CMD ["/app/src/main"]