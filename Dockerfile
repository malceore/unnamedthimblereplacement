FROM golang:1.11
LABEL maintainer="Brandon Wood	<malceore@gmail.com>"
WORKDIR /home/b/Code/gocode/src/github.com/malceore/thimble_replacement/
COPY . /home/b/Code/gocode/src/github.com/malceore/thimble_replacement/
EXPOSE 9191

RUN go version
RUN set -x && \
    go get github.com/golang/dep/cmd/dep && \
    go get github.com/gorilla/mux && \
    go get github.com/gorilla/securecookie
    #dep init

#RUN "curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh"
#RUN go get -d -v ./...
#RUN go install -v ./...
#RUN dep init
#go get -d -v ./...
RUN go build -o main
CMD ["./main"]