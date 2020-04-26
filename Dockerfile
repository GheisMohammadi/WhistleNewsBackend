FROM golang:latest

RUN go get github.com/golang/dep/cmd/dep
RUN go get github.com/posener/wstest
RUN go get github.com/oxequa/realize
RUN go get github.com/stretchr/testify

WORKDIR /go/src/
RUN mkdir -p WhistleNews
WORKDIR /go/src/WhistleNews

COPY . .

RUN dep ensure -update

RUN go build
RUN go install .

#RUN go get -t -v ./...
#RUN go install .
#RUN go-wrapper download
#RUN go-wrapper install

EXPOSE 4040