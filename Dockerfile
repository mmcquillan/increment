FROM golang:1.10.0-alpine3.7
WORKDIR /go/src/increment
COPY increment.go .
ENV GOBIN=/go/bin
RUN go install increment.go
CMD ["increment"]

