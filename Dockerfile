FROM golang:alpine as build

RUN apk update && apk add --no-cache git 

WORKDIR $GOPATH/src/github.com/Nerzal/CardsOfBinokee-Server
COPY . .

RUN go get -d -v; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app

FROM scratch

COPY --from=builder /go/bin/app bin/app

EXPOSE 995

CMD ["/bin/app"]