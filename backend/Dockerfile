FROM golang:alpine as builder

ENV SRC=/go/src/github.com/ComputePractice2018/microblog/backend
RUN mkdir -p '${SRC}'
ADD . ${SRC}
RUN ls ${SRC} 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/microblog github.com/ComputePractice2018/microblog/backend 
 

FROM alpine:3.7

WORKDIR /opt 
COPY --from=bulilder /go/bin/microblog /opt/microblog

EXPOSE 8080

ENV NAME "Irina"
ENTRYPOINT ["/go/bin/microblog"]