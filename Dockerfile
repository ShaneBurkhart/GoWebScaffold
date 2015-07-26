FROM golang:1.4

MAINTAINER Shane Burkhart <shaneburkhart@gmail.com>

RUN apt-get update \
    && apt-get -y install npm \
    && apt-get -y install nodejs

RUN ln -s /usr/bin/nodejs /usr/bin/node

# TODO remove this into a dev env.
RUN go get github.com/codegangsta/gin
RUN go get bitbucket.org/liamstask/goose/cmd/goose

ADD . /go/src/github.com/ShaneBurkhart/PlanSource

ENV GOPATH /go/src/github.com/ShaneBurkhart/PlanSource/vendor:${GOPATH}

WORKDIR /go/src/github.com/ShaneBurkhart/PlanSource

RUN npm install
# Build JS
#RUN ./node_modules/.bin/gulp

EXPOSE 3000

CMD ["go", "run", "main.go"]
