FROM google/golang

RUN cd /goroot/src/ && GOOS=linux GOARCH=386 ./make.bash --no-clean

WORKDIR /gopath/src/uuid5

ADD . /gopath/src/uuid

RUN go get
