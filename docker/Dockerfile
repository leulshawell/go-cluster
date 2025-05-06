FROM golang:1.24

WORKDIR /

RUN git clone https://github.com/leulshawell/go-cluster

RUN ls -l

WORKDIR /go-cluster/go-cluster

RUN ls -l 

RUN go mod tidy

CMD [ "go", "run", "./test/main.go"]