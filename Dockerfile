FROM golang:alpine as build

RUN apk update && apk add --no-cache git 

RUN go get -u -v github.com/Nerzal/CardsOfBinokee-Server/cmd/service
WORKDIR  /go/src/github.com/Nerzal/CardsOfBinokee-Server/cmd/service

RUN go get -d -v; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app

FROM scratch

COPY --from=build /go/src/github.com/Nerzal/CardsOfBinokee-Server/cmd/service /bin/app

EXPOSE 995

CMD ["/bin/app"]