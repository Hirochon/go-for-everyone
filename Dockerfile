FROM golang:1.18.2-alpine3.15
WORKDIR /go/src/work
COPY . /go/src/work/

## gccとmusl-dev は runtime/cgo で必要
RUN apk update && \
    apk add --no-cache git gcc musl-dev \
    go install -v golang.org/x/tools/cmd/godoc

EXPOSE 2300
EXPOSE 6060
