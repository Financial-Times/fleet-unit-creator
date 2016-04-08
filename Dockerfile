FROM alpine

RUN apk --update add go git gcc linux-headers libc-dev bash \
  && export GOPATH=/gopath \
  && mkdir $GOPATH \
  && go get github.com/Financial-Times/fleet-unit-creator 

CMD export GOPATH=/gopath \
  && cd $GOPATH/src/github.com/coreos/fleet \
  && git checkout $FLEET_REVISION \
  && go build -a -o /fleet-unit-creator github.com/Financial-Times/fleet-unit-creator \
  && /fleet-unit-creator -rootURI=$ROOT_URI