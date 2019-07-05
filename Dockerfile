FROM alpine:latest
RUN apk add --no-cache --update \
     && apk add go ca-certificates libssh2 libcurl expat pcre git make git musl-dev

ADD main.go main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /root/hget .

FROM alpine:latest  
#RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /root/hget /usr/bin/hget
CMD ["/usr/bin/hget"]  
