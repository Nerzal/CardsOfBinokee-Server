FROM golang:alpine as build

RUN apk update && apk add --no-cache git 

RUN go get -u -v github.com/Nerzal/CardsOfBinokee-Server/cmd/service
WORKDIR  /go/src/github.com/Nerzal/CardsOfBinokee-Server/cmd/service

RUN go get -d -v; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app

FROM scratch

ENV MONGODB_SERVER localhost
ENV MONGODB_DATABASE cardsofbinokee
ENV MONGODB_SSL false

# MONGODB_USERNAME 
# MONGODB_PASSWORD 
# MONGODB_REPLICASET_NAME 

COPY --from=build /go/src/github.com/Nerzal/CardsOfBinokee-Server/cmd/service/app /bin/app

EXPOSE 995
EXPOSE 8090

CMD ["/bin/app"]