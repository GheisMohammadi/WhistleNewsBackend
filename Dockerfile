FROM golang:latest

#install Whisle News Backend
RUN go get github.com/golang/dep/cmd/dep
RUN go get github.com/posener/wstest
RUN go get github.com/oxequa/realize
RUN go get github.com/stretchr/testify

WORKDIR /go/src/
RUN mkdir -p WhistleNewsBackend
WORKDIR /go/src/WhistleNewsBackend

COPY . .

RUN dep ensure -update

RUN go build -o whistlenewsservice
RUN go install .

EXPOSE 3085

RUN chmod 777 whistlenewsservice
#RUN ./whistlenewsservice

CMD ["/go/src/WhistleNewsBackend/whistlenewsservice"]