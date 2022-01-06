FROM golang:1.16.12-alpine3.15

RUN apk add --no-cache bash coreutils make git
RUN go get -u golang.org/x/lint/golint
RUN mkdir -p /go/src/workspace/shoppinglist/log

WORKDIR /go/src/workspace/shoppinglist
COPY . .
RUN make build
EXPOSE 9595

ENTRYPOINT [ "./shopping_list" ]