FROM golang:1.12
WORKDIR /gin-web-site
COPY . .
RUN export GOROOT=/usr/local/go
RUN export GOPATH=$HOME/go
RUN export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
RUN export PATH=$PATH:/usr/local/go/bin
RUN export PATH=$PATH:/go/bin
RUN go mod download
# RUN go install ./cmd/gin-web-site/main.go
RUN go build cmd/gin-web-site/main.go
# CMD ["go", "run", "cmd/gin-web-site/main.go"]
CMD ["./main", "runserver"]
EXPOSE 8081

# FROM redis
# COPY redis.conf /usr/local/etc/redis/redis.conf
# CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]
