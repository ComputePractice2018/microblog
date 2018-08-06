FROM golang:alpine as builder

RUN apk add --no-cache git
RUN go get -u github.com/jinzhu/gorm 
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gorilla/mux

ENV SRC=/go/src/github.com/ComputePractice2018/microblog/backend/
RUN mkdir -p '${SRC}'
ADD . ${SRC}


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/microblog github.com/ComputePractice2018/microblog/backend 
 

FROM alpine:3.7

WORKDIR /opt 
COPY --from=builder /go/bin/microblog /opt/microblog

EXPOSE 8080

ENV NAME "Irina"
ENTRYPOINT /opt/microblog -name ${NAME}