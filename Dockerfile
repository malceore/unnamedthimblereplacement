#build stage
#FROM golang:1.11 AS build-env
FROM golang:alpine AS build-env
LABEL maintainer="Brandon Wood<malceore@gmail.com>"
RUN apk update && apk add --no-cache git
ADD . /src
#WORKDIR /home/b/Code/gocode/src/github.com/malceore/thimble_replacement/
#COPY . /home/b/Code/gocode/src/github.com/malceore/thimble_replacement/
#RUN go version && go test
RUN set -x && \
    go get github.com/gorilla/mux && \
    go get github.com/gorilla/securecookie &&\
    go get -u github.com/lib/pq
RUN ls src
RUN cd /src && go build -o main

# final stage
FROM alpine AS app
WORKDIR /app
RUN mkdir -p /app/res
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=build-env /src/bin/main /app/
COPY --from=build-env /src/res /app/res
EXPOSE 9191
CMD ["./main"]
#ENTRYPOINT /app/./main
