FROM golang:1.12.9-alpine AS build_base

RUN apk add --update gcc g++ git

WORKDIR /go/src/github.com/example-chaincode

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base AS server_builder

COPY . .

RUN go mod vendor

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' .

FROM alpine AS weaviate

RUN apk add bash

COPY --from=server_builder /go/bin /bin/chaincode

CMD /bin/chaincode/helloworld-chaincode