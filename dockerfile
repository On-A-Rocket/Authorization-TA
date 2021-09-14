FROM golang:alpine AS stage
WORKDIR /go/src/github.com/On-A-Rocket/Authorization-TA
COPY . .
RUN apk update && \
  apk upgrade && \
  apk add --update-cache tzdata 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o main main.go

FROM scratch
COPY --from=stage /go/src/github.com/On-A-Rocket/Authorization-TA/main /main
COPY --from=stage /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
ENV TZ=Asia/Seoul \
    ZONEINFO=/zoneinfo.zip
EXPOSE 1323
CMD [ "/main" ]
