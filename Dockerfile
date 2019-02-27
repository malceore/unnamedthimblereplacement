FROM golang:alpine
RUN apk update
WORKDIR /home/b/Code/gocode/src/github.com/malceore/thimble_replacement
COPY . .
RUN go build -o bin/main
ENTRYPOINT ["bin/./hello"]